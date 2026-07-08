package bindings

import (
	"context"
	"fmt"

	"github.com/thani-sh/provar/libs/domain"
	"github.com/thani-sh/provar/libs/engine"
)

// Run is the desktop app binding for executing committed Lua. The editor's
// Compile button writes the .test.lua; the Run button (and the CI-shaped
// "Run" path that follows) reads it back and drives the engine. Running
// requires no LLM session — the compiled script is already there.
//
// The binding embeds baseJobs so the per-job lifecycle (event forwarding,
// cancel, cleanup) lives in one place. Run itself is just the
// "load → run" glue plus the Wails-exposed methods.
type Run struct {
	baseJobs
}

// Start loads the project config + compiled Lua for relPath and launches
// an engine.Runner.Run. The job ID returned is stable for the run's
// lifetime; cancelRun uses it to find the matching engine Job and ask it
// to stop.
//
// Errors are returned immediately so the toolbar can show a toast:
//   - project config missing → user opened the project before .provar/config.yml exists
//   - .test.yml unreadable → the file is gone or permission denied
//   - .test.lua missing or stale → user hasn't compiled yet (or its mod-time predates the .yml)
//   - engine Run returned an error before the goroutine started → likely a malformed Lua script
func (r *Run) Start(projectDir, relPath string, headless bool, upTo string) (string, error) {
	project, err := domain.LoadProject(projectDir)
	if err != nil {
		return "", fmt.Errorf("load project: %w", err)
	}
	actions, err := domain.ParseFile(projectDir, relPath)
	if err != nil {
		return "", fmt.Errorf("parse test file: %w", err)
	}
	luaCode, ok := domain.LoadCompiledLua(projectDir, relPath)
	if !ok {
		return "", fmt.Errorf("compiled .test.lua not found or out of date; run `provar compile` first")
	}

	ctx, cancel := context.WithCancel(r.Ctx)
	runner := engine.NewRunner()
	job, err := runner.Run(ctx, actions, luaCode, engine.RunOptions{
		Headless: headless,
		Vars:     project.Vars,
		UpTo:     upTo,
		Browser:  project.Browser,
	})
	if err != nil {
		cancel()
		return "", fmt.Errorf("start run: %w", err)
	}

	// Runner.Run defers session.Close inside its goroutine; we don't need
	// a cleanup hook here.
	return r.trackJob(job, cancel, nil), nil
}

// Cancel stops an in-flight run by id. Returns nil when the id is unknown
// (the job already finished); the toolbar treats that as a no-op success
// because the run is over either way.
func (r *Run) Cancel(jobID string) error {
	return r.baseJobs.Cancel(jobID)
}
