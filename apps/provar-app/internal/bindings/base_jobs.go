package bindings

import (
	"fmt"
	"sync"

	"github.com/thani-sh/provar/libs/domain"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// EventName is the Wails event-name used by both Run and Compile. Binding
// events use a single shared namespace because the consumer (the frontend)
// already filters by jobId; we don't ship a separate name per type.
const EventName = "job:event"

// baseJobs is the lifecycle scaffold shared by Run and Compile. Each
// concrete binding embeds it (see the per-binding files for the wrapping
// Start/Cancel surface) so the event-forwarding, job map, and cleanup
// dance aren't duplicated.
//
// A binding tracks every job it has started in `jobs` keyed by the engine's
// domain.Job.ID so Cancel can find the right one without callers passing
// a per-binding handle. The map is the only state the binding owns across
// calls; no global locks required.
type baseJobs struct {
	BaseBinding

	mu   sync.Mutex
	jobs map[string]*jobHandle
}

// jobHandle pairs the engine job with the context that cancellation will
// cancel. The cancel func is the same context used to start the engine
// run, so calling it races cleanly with the engine's own stop signal.
type jobHandle struct {
	job    *domain.Job
	cancel func()
	// done is closed when the emit loop exits; bindings don't read it
	// today, but it lets a future RPC wait synchronously for "job
	// finished" without polling.
	done chan struct{}
}

// trackJob registers an engine job and forwards its events to the
// frontend. onCleanup, when non-nil, runs after the job's Subscribe
// channel closes (i.e. after the engine has emitted run-finished /
// compile-finished and called Job.Close). Compile uses this to release
// the browser session it owns; Run lets the engine's defer close it.
//
// trackJob returns the engine's job id — callers pass that back to the
// frontend so the editor store can scope its subscription with forJob().
func (b *baseJobs) trackJob(job *domain.Job, cancel func(), onCleanup func()) string {
	h := &jobHandle{
		job:    job,
		cancel: cancel,
		done:   make(chan struct{}),
	}
	b.mu.Lock()
	if b.jobs == nil {
		b.jobs = map[string]*jobHandle{}
	}
	b.jobs[job.ID] = h
	b.mu.Unlock()

	go b.emitLoop(h, onCleanup)
	return job.ID
}

// Cancel stops a tracked job by id. Returns an error if no such job is
// tracked — typically because the id is unknown to this process, or
// because the job has already completed and been swept from the map.
// Either way the frontend treats the call as a no-op success.
func (b *baseJobs) Cancel(jobID string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	h, ok := b.jobs[jobID]
	if !ok {
		return fmt.Errorf("no job with id %s", jobID)
	}
	h.cancel()
	// Belt-and-braces: the engine's runWaitLoop checks JobStopped as well
	// as ctx.Done(), so without the Stop() a future action could miss the
	// cancellation. Setting both gives the fastest exit.
	h.job.Stop()
	return nil
}

// emitLoop consumes the job's Subscribe channel and re-emits each event
// through Wails's runtime.EventsEmit under EventName. The payload is a
// single object so the frontend's `for await (const ev of subscribe(name))`
// receives it as one argument.
//
// On loop exit (channel closed by Job.Close or job completed naturally),
// the registration entry is removed and any provided cleanup runs.
func (b *baseJobs) emitLoop(h *jobHandle, onCleanup func()) {
	defer close(h.done)
	defer func() {
		if onCleanup != nil {
			onCleanup()
		}
	}()
	defer func() {
		b.mu.Lock()
		delete(b.jobs, h.job.ID)
		b.mu.Unlock()
	}()

	for ev := range h.job.Subscribe() {
		runtime.EventsEmit(b.Ctx, EventName, map[string]any{
			"jobId": h.job.ID,
			"type":  ev.Type,
			"data":  ev.Data,
		})
	}
}
