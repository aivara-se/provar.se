import { forJob } from './events';
import { executionStore } from '../stores/execution-store.svelte';

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

          const payload = event.data as Record<string, unknown> | undefined;
          const actionId =
            (payload?.actionId as string | undefined) ??
            (payload?.ActionID as string | undefined) ??
            (payload?.actionID as string | undefined);
          const name = (payload?.Name as string | undefined) ?? (payload?.name as string | undefined);
          const errText =
            (payload?.Error as string | undefined) ??
            (payload?.error as string | undefined) ??
            (typeof payload === 'string' ? payload : undefined);
          const duration =
            (payload?.Duration as string | undefined) ?? (payload?.duration as string | undefined);

          switch (event.type) {
            case 'run-started':
              executionStore.addLog({
                level: 'info',
                category: 'run',
                message: 'Test execution engine started',
              });
              break;

            case 'compile-started':
              executionStore.addLog({
                level: 'info',
                category: 'compile',
                message: 'Test compiler started LLM authoring loop',
              });
              break;

            case 'task-started':
              if (actionId) {
                callbacks.onTaskStarted?.(actionId);
                executionStore.addLog({
                  level: 'step',
                  category: 'run',
                  actionId,
                  message: `Executing action step ${name ? `"${name}" ` : ''}[${actionId}]`,
                });
              }
              break;

            case 'task-finished':
              if (actionId) {
                callbacks.onTaskFinished?.(actionId);
                executionStore.addLog({
                  level: 'success',
                  category: 'run',
                  actionId,
                  message: `Action step [${actionId}] completed successfully`,
                });
              }
              break;

            case 'task-failed':
              if (actionId) {
                callbacks.onTaskFailed?.(actionId);
                executionStore.addLog({
                  level: 'error',
                  category: 'run',
                  actionId,
                  message: `Action step [${actionId}] failed`,
                  details: errText,
                });
              }
              break;

            case 'run-finished':
              callbacks.onRunFinished?.();
              if (errText) {
                executionStore.addLog({
                  level: 'error',
                  category: 'run',
                  message: `Test execution failed: ${errText}`,
                  details: errText,
                });
              } else {
                executionStore.addLog({
                  level: 'success',
                  category: 'run',
                  message: `Test execution finished ${duration ? `in ${duration}` : 'successfully'}`,
                });
              }
              cancel();
              break;

            case 'action-started':
              if (actionId) {
                callbacks.onActionStarted?.(actionId);
                executionStore.addLog({
                  level: 'step',
                  category: 'compile',
                  actionId,
                  message: `Authoring Lua for action ${name ? `"${name}" ` : ''}[${actionId}]...`,
                });
              }
              break;

            case 'action-finished':
              if (actionId) {
                callbacks.onActionFinished?.(actionId);
                executionStore.addLog({
                  level: 'success',
                  category: 'compile',
                  actionId,
                  message: `Action node [${actionId}] compiled successfully`,
                });
              }
              break;

            case 'action-failed':
              if (actionId) {
                callbacks.onActionFailed?.(actionId);
                executionStore.addLog({
                  level: 'error',
                  category: 'compile',
                  actionId,
                  message: `Action node [${actionId}] failed compilation`,
                  details: errText,
                });
              }
              break;

            case 'visual-comparison-triggered':
              if (actionId) {
                executionStore.addLog({
                  level: 'info',
                  category: 'browser',
                  actionId,
                  message: `Visual comparison assertion triggered for action node [${actionId}]`,
                });
              }
              break;

            case 'compile-finished':
              callbacks.onCompileFinished?.();
              if (errText) {
                executionStore.addLog({
                  level: 'error',
                  category: 'compile',
                  message: `Compilation failed: ${errText}`,
                  details: errText,
                });
              } else {
                executionStore.addLog({
                  level: 'success',
                  category: 'compile',
                  message: `Compilation completed ${duration ? `in ${duration}` : 'successfully'}`,
                });
              }
              cancel();
              break;

            case 'log':
              executionStore.addLog({
                level: (payload?.level as any) || 'info',
                category: (payload?.category as any) || 'system',
                actionId,
                message: (payload?.message as string) || String(event.data),
                details: payload?.details as string | undefined,
              });
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
