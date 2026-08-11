// Wails event bus wrapped as async iterators. Bindings push events via
// runtime.EventsEmit (Go side) and the frontend consumes them through here.
//
// The wrapper exists so call sites read as ordinary async iteration rather
// than a callback juggling its own cleanup. Each call to `subscribe` opens
// its own EventsOn subscription; the subscriber is released when the loop
// exits, whether it ran to completion, broke, or threw — the off function
// returned by EventsOn is invoked in the generator's `finally`.
//
// `forJob` filters the same event stream by `jobId`, which is how the editor
// scopes its per-job subscription to the Compile/Run it started. Multiple
// jobs can be in flight if the user races the toolbar; the filter keeps
// each for-await loop reading only its own job's events.

import { EventsOn } from '../../wailsjs/runtime/runtime';

type ReleaseFn = () => void;

/**
 * subscribe yields every payload the bindings publish on `name`.
 *
 * Each call returns a fresh iterator that registers its own EventsOn
 * listener. The listener is released when the loop terminates for any
 * reason (early break, throw, or normal completion). Payloads are the
 * first argument EventsEmit received — multi-argument emits are not
 * supported here because the editor's bindings always pass one object.
 */
export async function* subscribe<T = unknown>(name: string): AsyncGenerator<T, void, void> {
  const queue: T[] = [];
  const waiting: Array<(r: IteratorResult<T, void>) => void> = [];
  const off = register<T>(name, queue, waiting);
  try {
    while (true) {
      if (queue.length > 0) {
        yield queue.shift()!;
        continue;
      }
      const next = await new Promise<IteratorResult<T, void>>((resolve) => {
        waiting.push(resolve);
      });
      if (next.done) return;
      yield next.value;
    }
  } finally {
    off();
  }
}

/**
 * forJob is subscribe + a jobId filter, for events whose payload carries
 * a `jobId` field (all Compile/Run events do). The filter is applied
 * inside the generator so dropped events are discarded immediately and
 * never reach the consumer.
 */
export async function* forJob<T extends { jobId: string }>(
  jobId: string,
  name: string,
): AsyncGenerator<T, void, void> {
  for await (const event of subscribe<T>(name)) {
    if (event.jobId === jobId) yield event;
  }
}

function register<T>(name: string, queue: T[], waiting: Array<(r: IteratorResult<T, void>) => void>): ReleaseFn {
  // Wails's EventsOn returns an off function so multiple subscribers can
  // coexist. We use that handle rather than EventsOff(name) because the
  // latter nukes every subscriber on `name`, not just ours.
  let released = false;
  const off = EventsOn(name, (payload: T) => {
    if (released) return;
    const next = waiting.shift();
    if (next) next({ value: payload, done: false });
    else queue.push(payload);
  });
  return () => {
    if (released) return;
    released = true;
    off();
    // Drain any pending waits with a done so a consumer that's blocked on
    // the next event after break returns promptly instead of hanging.
    while (waiting.length > 0) {
      waiting.shift()!({ value: undefined as never, done: true });
    }
  };
}
