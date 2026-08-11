import { HistoryService } from '../services/history-service';

/**
 * HistoryStore owns the desktop app's recent-projects list. Persists
 * via HistoryService to ~/.provar/history.yml.
 */
class HistoryStore {
  recent = $state<string[]>([]);
  loaded = $state(false);

  /**
   * load reads the on-disk history into the store. Idempotent.
   */
  async load() {
    if (this.loaded) return;
    this.loaded = true;
    this.recent = await HistoryService.getRecent();
  }

  /**
   * add prepends path optimistically, then persists. On error, reverts
   * to the prior list and re-throws.
   */
  async add(path: string): Promise<void> {
    const prior = this.recent;
    this.recent = [path, ...prior.filter((p) => p !== path)].slice(0, 10);
    try {
      await HistoryService.addRecent(path);
    } catch (err) {
      this.recent = prior;
      throw err;
    }
  }
}

export const historyStore = new HistoryStore();