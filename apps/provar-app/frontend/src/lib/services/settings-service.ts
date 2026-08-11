import { Project } from './bindings';
import { domain } from '../../../wailsjs/go/models';
import { applicationStore } from '../stores/application-store.svelte';

export class SettingsService {
  static async getHome(): Promise<string> {
    try {
      return await Project.Home();
    } catch (e) {
      console.error('SettingsService: Home lookup failed:', e);
      throw e;
    }
  }

  static async getSettings(): Promise<domain.Settings> {
    try {
      const settings = await Project.Settings();
      return settings ?? new domain.Settings();
    } catch (e) {
      console.error('SettingsService: getSettings failed:', e);
      throw e;
    }
  }

  static async saveSettings(settings: domain.Settings): Promise<void> {
    try {
      await Project.SaveSettings(settings);
      applicationStore.showToast('info', 'Settings saved successfully');
    } catch (e) {
      const msg = e instanceof Error ? e.message : String(e);
      applicationStore.showToast('error', `Could not save settings: ${msg}`);
      throw e;
    }
  }

  static async validateSettings(): Promise<void> {
    const err = await Project.ValidateSettings();
    if (err !== null) {
      throw new Error(String(err));
    }
  }
}
