import { forJob } from './events';

export interface JobCallbacks {
  onTaskStarted?: (actionId: string) => void;
  onTaskFinished?: (actionId: string) => void;
  onTaskFailed?: (actionId: string) => void;
  onRunFinished?: () => void;

  onActionStarted?: (actionId: string) => void;
  onActionFinished?: (actionId: string) => void;
  onActionFailed?: (actionId: string) => void;
  onCompileFinished?: () => void;
}

export class JobStreamManager {
  private activeJobs = new Map<string, { cancel: () => void }>();

  subscribeJob(jobId: string, callbacks: JobCallbacks): () => void {
    let cancelled = false;

    const cancel = () => {
      cancelled = true;
      this.activeJobs.delete(jobId);
    };

    this.activeJobs.set(jobId, { cancel });

    void (async () => {
      try {
        for await (const event of forJob<{ jobId: string; type: string; data?: unknown }>(jobId, 'job:event')) {
          if (cancelled) break;

          const payload = event.data as { actionId?: string; ActionID?: string } | undefined;
          const actionId = payload?.actionId ?? payload?.ActionID;

          switch (event.type) {
            case 'task-started':
              if (actionId) callbacks.onTaskStarted?.(actionId);
              break;
            case 'task-finished':
              if (actionId) callbacks.onTaskFinished?.(actionId);
              break;
            case 'task-failed':
              if (actionId) callbacks.onTaskFailed?.(actionId);
              break;
            case 'run-finished':
              callbacks.onRunFinished?.();
              cancel();
              break;

            case 'action-started':
              if (actionId) callbacks.onActionStarted?.(actionId);
              break;
            case 'action-finished':
              if (actionId) callbacks.onActionFinished?.(actionId);
              break;
            case 'action-failed':
              if (actionId) callbacks.onActionFailed?.(actionId);
              break;
            case 'compile-finished':
              callbacks.onCompileFinished?.();
              cancel();
              break;
          }
        }
      } catch (e) {
        console.error(`JobStreamManager error processing stream for job ${jobId}:`, e);
      } finally {
        this.activeJobs.delete(jobId);
      }
    })();

    return cancel;
  }

  stopJob(jobId: string): void {
    const job = this.activeJobs.get(jobId);
    if (job) {
      job.cancel();
    }
  }

  stopAll(): void {
    for (const job of this.activeJobs.values()) {
      job.cancel();
    }
    this.activeJobs.clear();
  }
}

export const jobStreamManager = new JobStreamManager();
