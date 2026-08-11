import { File } from './bindings';
import type { TestFileView } from '../domain/types';

export class FileService {
  static async createFile(absPath: string): Promise<void> {
    try {
      await File.CreateFile(absPath);
    } catch (e) {
      console.error('FileService: createFile failed:', e);
      throw e;
    }
  }

  static async createDirectory(absPath: string): Promise<void> {
    try {
      await File.CreateDirectory(absPath);
    } catch (e) {
      console.error('FileService: createDirectory failed:', e);
      throw e;
    }
  }

  static async deletePath(absPath: string): Promise<void> {
    try {
      await File.DeletePath(absPath);
    } catch (e) {
      console.error('FileService: deletePath failed:', e);
      throw e;
    }
  }

  static async readTestFile(projectPath: string, relPath: string): Promise<TestFileView> {
    try {
      const view = await File.ReadTestFile(projectPath, relPath);
      return {
        graph: view.graph,
        order: view.order ?? [],
      };
    } catch (e) {
      console.error('FileService: readTestFile failed:', e);
      throw e;
    }
  }

  static async writeTestFile(
    projectPath: string,
    relPath: string,
    view: TestFileView,
  ): Promise<void> {
    try {
      await File.WriteTestFile(projectPath, relPath, view as any);
    } catch (e) {
      console.error('FileService: writeTestFile failed:', e);
      throw e;
    }
  }

  static async validate(
    projectPath: string = '',
    relPath: string = '',
    view?: TestFileView,
  ): Promise<import('../domain/types').DiagnosticReport> {
    try {
      return await File.Validate(projectPath, relPath, view as any);
    } catch (e) {
      console.error('FileService: validate failed:', e);
      return { isValid: true, errors: [], warnings: [] };
    }
  }
}
