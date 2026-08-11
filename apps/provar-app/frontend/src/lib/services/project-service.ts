import { Dialog, Project, Config, File, Watcher } from './bindings';

export class ProjectService {
  static async selectProjectDialog(): Promise<string | null> {
    try {
      const path = await Dialog.SelectProject();
      return path || null;
    } catch (e) {
      console.error('ProjectService: selectProjectDialog failed:', e);
      throw e;
    }
  }

  static async createSampleProject(path: string): Promise<void> {
    try {
      await Project.CreateSampleProject(path);
    } catch (e) {
      console.error('ProjectService: createSampleProject failed:', e);
      throw e;
    }
  }

  static async loadConfig(path: string): Promise<Record<string, unknown> | null> {
    try {
      return await Config.LoadConfig(path);
    } catch (e) {
      console.error('ProjectService: loadConfig failed:', e);
      return null;
    }
  }

  static async saveConfig(path: string, config: Record<string, unknown>): Promise<void> {
    try {
      await Config.SaveConfig(path, config);
    } catch (e) {
      console.error('ProjectService: saveConfig failed:', e);
      throw e;
    }
  }

  static async listTests(path: string): Promise<string[]> {
    try {
      return await File.ListTests(path);
    } catch (e) {
      console.error('ProjectService: listTests failed:', e);
      return [];
    }
  }

  static async watchProject(path: string): Promise<void> {
    try {
      await Watcher.Watch(path);
    } catch (e) {
      console.error('ProjectService: watchProject failed:', e);
      throw e;
    }
  }
}
