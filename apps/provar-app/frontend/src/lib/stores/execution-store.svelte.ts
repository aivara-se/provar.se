import { ExecutionService } from '../services/execution-service';
import { jobStreamManager } from '../services/job-stream-manager';
import { applicationStore } from './application-store.svelte';

export type RunState = 'idle' | 'running' | 'success' | 'failed';
export type CompileState = 'idle' | 'compiling' | 'compiled' | 'failed';

export interface LogEntry {
  id: string;
  timestamp: string;
  level: 'info' | 'success' | 'warn' | 'error' | 'step';
  message: string;
  category?: 'compile' | 'run' | 'browser' | 'system';
  actionId?: string;
  details?: string;
}

class ExecutionStore {
  isRunning = $state(false);
  activeRunId = $state<string | null>(null);
  isCompiling = $state(false);
  activeCompileId = $state<string | null>(null);

  taskStates = $state<Record<string, RunState>>({});
  compileStates = $state<Record<string, CompileState>>({});
  logs = $state<LogEntry[]>([]);

  clearStates() {
    this.taskStates = {};
    this.compileStates = {};
  }

  addLog(entry: Omit<LogEntry, 'id' | 'timestamp'> & { timestamp?: string }) {
    const now = new Date();
    const ts =
      entry.timestamp ??
      `${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}:${now.getSeconds().toString().padStart(2, '0')}.${now.getMilliseconds().toString().padStart(3, '0')}`;
    const newLog: LogEntry = {
      id: `${Date.now()}-${Math.random().toString(36).substring(2, 7)}`,
      timestamp: ts,
      ...entry,
    };
    this.logs = [...this.logs, newLog];
  }

  clearLogs() {
    this.logs = [];
  }

  async runCurrent(projectPath: string, testPath: string, upTo = ''): Promise<void> {
    if (this.isRunning) return;
    this.taskStates = {};
    applicationStore.openConsole();
    this.addLog({
      level: 'info',
      category: 'run',
      message: `Starting test execution for ${testPath}${upTo ? ` (up to ${upTo})` : ''}...`,
    });

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
      this.addLog({
        level: 'error',
        category: 'run',
        message: `Could not start run: ${msg}`,
      });
      this.isRunning = false;
      this.activeRunId = null;
    }
  }

  async stopRun(): Promise<void> {
    const jobId = this.activeRunId;
    if (!jobId) return;
    try {
      this.addLog({
        level: 'warn',
        category: 'run',
        message: 'Stopping test execution...',
      });
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
    applicationStore.openConsole();
    this.addLog({
      level: 'info',
      category: 'compile',
      message: `Starting test compilation for ${testPath}...`,
    });

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
      this.addLog({
        level: 'error',
        category: 'compile',
        message: `Could not start compile: ${msg}`,
      });
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
