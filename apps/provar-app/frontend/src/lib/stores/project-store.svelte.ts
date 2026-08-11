import { ProjectService } from '../services/project-service';
import { historyStore } from './history-store.svelte';

/**
 * ProjectStore owns the path, config, and test-file list of the active
 * project. Recent-projects persistence lives in historyStore.
 */
class ProjectStore {
  path = $state<string | null>(null);
  config = $state<Record<string, unknown> | null>(null);
  tests = $state<string[]>([]);

  setPath(path: string | null) {
    this.path = path;
    if (path === null) {
      this.tests = [];
      this.config = null;
    }
  }

  async refreshTests() {
    if (!this.path) {
      this.tests = [];
      return;
    }
    this.tests = await ProjectService.listTests(this.path);
  }

  async openProject(path: string) {
    this.setPath(path);
    await Promise.all([this.refreshTests(), this.loadConfig()]);
    try {
      await historyStore.add(path);
    } catch (e) {
      console.error('ProjectStore: history add failed:', e);
    }
  }

  async loadConfig() {
    if (!this.path) return;
    this.config = await ProjectService.loadConfig(this.path);
  }

  async saveConfig(next: Record<string, unknown>) {
    if (!this.path) return;
    await ProjectService.saveConfig(this.path, next);
    this.config = next;
  }
}

export const projectStore = new ProjectStore();