import { History } from './bindings';

export class HistoryService {
  static async getRecent(): Promise<string[]> {
    try {
      return await History.Recent();
    } catch (e) {
      console.error('HistoryService: getRecent failed:', e);
      return [];
    }
  }

  static async addRecent(path: string): Promise<void> {
    try {
      await History.Add(path);
    } catch (e) {
      console.error('HistoryService: addRecent failed:', e);
      throw e;
    }
  }
}
