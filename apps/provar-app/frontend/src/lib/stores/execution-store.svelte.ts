import { ExecutionService } from '../services/execution-service';
import { jobStreamManager } from '../services/job-stream-manager';
import { applicationStore } from './application-store.svelte';

export type RunState = 'idle' | 'running' | 'success' | 'failed';
export type CompileState = 'idle' | 'compiling' | 'compiled' | 'failed';

class ExecutionStore {
  isRunning = $state(false);
  activeRunId = $state<string | null>(null);
  isCompiling = $state(false);
  activeCompileId = $state<string | null>(null);

  taskStates = $state<Record<string, RunState>>({});
  compileStates = $state<Record<string, CompileState>>({});

  clearStates() {
    this.taskStates = {};
    this.compileStates = {};
  }

  async runCurrent(projectPath: string, testPath: string, upTo = ''): Promise<void> {
    if (this.isRunning) return;
    this.taskStates = {};
    try {
      const jobId = await ExecutionService.startRun(projectPath, testPath, true, upTo);
      this.activeRunId = jobId;
      this.isRunning = true;

      jobStreamManager.subscribeJob(jobId, {
        onTaskStarted: (actionId) => {
          this.taskStates = { ...this.taskStates, [actionId]: 'running' };
        },
        onTaskFinished: (actionId) => {
          this.taskStates = { ...this.taskStates, [actionId]: 'success' };
        },
        onTaskFailed: (actionId) => {
          this.taskStates = { ...this.taskStates, [actionId]: 'failed' };
        },
        onRunFinished: () => {
          this.isRunning = false;
          this.activeRunId = null;
        },
      });
    } catch (e) {
      const msg = e instanceof Error ? e.message : String(e);
      applicationStore.showToast('error', `Could not start run: ${msg}`);
      this.isRunning = false;
      this.activeRunId = null;
    }
  }

  async stopRun(): Promise<void> {
    const jobId = this.activeRunId;
    if (!jobId) return;
    try {
      jobStreamManager.stopJob(jobId);
      await ExecutionService.cancelRun(jobId);
    } catch {
      // Already finished or cancelled
    } finally {
      this.isRunning = false;
      this.activeRunId = null;
    }
  }

  async compileCurrent(
    projectPath: string,
    testPath: string,
    onSuccess?: () => Promise<void>,
  ): Promise<void> {
    if (this.isCompiling) return;
    this.compileStates = {};
    try {
      const jobId = await ExecutionService.startCompile(projectPath, testPath);
      this.activeCompileId = jobId;
      this.isCompiling = true;

      jobStreamManager.subscribeJob(jobId, {
        onActionStarted: (actionId) => {
          this.compileStates = { ...this.compileStates, [actionId]: 'compiling' };
        },
        onActionFinished: (actionId) => {
          this.compileStates = { ...this.compileStates, [actionId]: 'compiled' };
        },
        onActionFailed: (actionId) => {
          this.compileStates = { ...this.compileStates, [actionId]: 'failed' };
        },
        onCompileFinished: () => {
          this.isCompiling = false;
          this.activeCompileId = null;
          if (onSuccess) {
            void onSuccess();
          }
        },
      });
    } catch (e) {
      const msg = e instanceof Error ? e.message : String(e);
      applicationStore.showToast('error', `Could not start compile: ${msg}`);
      this.isCompiling = false;
      this.activeCompileId = null;
    }
  }

  async stopCompile(): Promise<void> {
    const jobId = this.activeCompileId;
    if (!jobId) return;
    try {
      jobStreamManager.stopJob(jobId);
      await ExecutionService.cancelCompile(jobId);
    } catch {
      // Already finished or cancelled
    } finally {
      this.isCompiling = false;
      this.activeCompileId = null;
    }
  }
}

export const executionStore = new ExecutionStore();
