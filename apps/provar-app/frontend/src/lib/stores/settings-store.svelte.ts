import { SettingsService } from '../services/settings-service';
import type { domain } from '../../../wailsjs/go/models';

/**
 * SettingsStore owns app-lifecycle state from the on-disk settings file:
 * the user's home directory, global settings, and the setup wizard flag.
 */
class SettingsStore {
  homeDir = $state('');
  showSetupWizard = $state(false);
  hasCheckedSetup = $state(false);
  settings = $state<domain.Settings | null>(null);

  /** load reads the on-disk home dir and validates settings. Idempotent. */
  async load() {
    if (this.hasCheckedSetup) return;
    this.hasCheckedSetup = true;
    try {
      this.homeDir = await SettingsService.getHome();
    } catch (err) {
      console.error('SettingsStore: home lookup failed:', err);
    }

    try {
      this.settings = await SettingsService.getSettings();
      await SettingsService.validateSettings();
      this.showSetupWizard = false;
    } catch (err) {
      console.error('SettingsStore: settings validation failed:', err);
      this.showSetupWizard = true;
    }
  }

  async updateSettings(next: domain.Settings) {
    await SettingsService.saveSettings(next);
    this.settings = next;
    this.showSetupWizard = false;
  }

  dismissSetupWizard() {
    this.showSetupWizard = false;
  }
}

export const settingsStore = new SettingsStore();