import { Run, Compile } from './bindings';

export class ExecutionService {
  static async startRun(
    projectPath: string,
    testPath: string,
    headless = true,
    upTo = '',
  ): Promise<string> {
    try {
      return await Run.Start(projectPath, testPath, headless, upTo);
    } catch (e) {
      console.error('ExecutionService: startRun failed:', e);
      throw e;
    }
  }

  static async cancelRun(jobId: string): Promise<void> {
    try {
      await Run.Cancel(jobId);
    } catch (e) {
      console.error('ExecutionService: cancelRun failed:', e);
      throw e;
    }
  }

  static async startCompile(projectPath: string, testPath: string): Promise<string> {
    try {
      return await Compile.Start(projectPath, testPath);
    } catch (e) {
      console.error('ExecutionService: startCompile failed:', e);
      throw e;
    }
  }

  static async cancelCompile(jobId: string): Promise<void> {
    try {
      await Compile.Cancel(jobId);
    } catch (e) {
      console.error('ExecutionService: cancelCompile failed:', e);
      throw e;
    }
  }
}
