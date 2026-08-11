import { subscribe } from './events';
import { applicationStore } from '../stores/application-store.svelte';

/**
 * AppCommands (Layer 3 Service per ADR-009) listens for Go-initiated
 * triggers and dispatches them to reactive state stores or services.
 */
export class AppCommands {
  /**
   * init registers listeners for all `app:*` Go event triggers.
   * Returns a cleanup function that releases event subscriptions when unmounted.
   */
  static init(): () => void {
    let cancelled = false;

    void (async () => {
      for await (const _ of subscribe('app:open-settings')) {
        if (cancelled) return;
        applicationStore.openSettingsModal();
      }
    })();

    void (async () => {
      for await (const _ of subscribe('app:toggle-test-explorer')) {
        if (cancelled) return;
        applicationStore.toggleSidebar();
      }
    })();

    void (async () => {
      for await (const _ of subscribe('app:toggle-console')) {
        if (cancelled) return;
        applicationStore.toggleConsole();
      }
    })();

    void (async () => {
      for await (const _ of subscribe('app:open-doctor')) {
        if (cancelled) return;
        applicationStore.openDoctorModal();
      }
    })();

    return () => {
      cancelled = true;
    };
  }
}
