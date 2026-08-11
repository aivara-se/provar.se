/**
 * ApplicationStore manages the visibility state of panels, sidebars, modals, and
 * transient messages. One state field per concern — no boolean explosion.
 *
 * Modals are dispatched by kind through `modalKind`. The active modal's
 * props live on `confirmProps` / `inputProps` so callers don't need to
 * thread the same payload through every consumer — AppModals reads
 * everything off the store and re-renders on every change.
 */

export type ModalKind = 'confirm' | 'input' | 'settings' | 'config' | null;

export interface Toast {
  id: number;
  kind: 'error' | 'info';
  message: string;
}

export interface ConfirmProps {
  title: string;
  message: string;
  onConfirm: () => void;
}

export interface InputProps {
  title: string;
  placeholder: string;
  onConfirm: (value: string) => void;
}

export type RightSidebarTab = 'project' | 'issues' | 'node' | 'console';

class ApplicationStore {
  isSidebarOpen = $state(true);
  isRightSidebarOpen = $state(false);
  rightSidebarTab = $state<RightSidebarTab>('project');
  modalKind = $state<ModalKind>(null);
  toast = $state<Toast | null>(null);

  // Modal payloads. The default values are no-op callbacks so the
  // renderer can render `confirmProps.title` unconditionally without
  // null-checking on every keystroke.
  confirmProps = $state<ConfirmProps>({
    title: '',
    message: '',
    onConfirm: () => {
      this.closeModal();
    },
  });
  inputProps = $state<InputProps>({
    title: '',
    placeholder: '',
    onConfirm: () => {
      this.closeModal();
    },
  });

  closeSidebars() {
    this.isSidebarOpen = false;
    this.isRightSidebarOpen = false;
  }

  toggleSidebar() {
    this.isSidebarOpen = !this.isSidebarOpen;
  }

  toggleRightSidebar() {
    this.isRightSidebarOpen = !this.isRightSidebarOpen;
  }

  toggleConsole() {
    if (this.isRightSidebarOpen && this.rightSidebarTab === 'console') {
      this.isRightSidebarOpen = false;
    } else {
      this.isRightSidebarOpen = true;
      this.rightSidebarTab = 'console';
    }
  }

  openConsole() {
    this.isRightSidebarOpen = true;
    this.rightSidebarTab = 'console';
  }

  closeConsole() {
    if (this.rightSidebarTab === 'console') {
      this.isRightSidebarOpen = false;
    }
  }

  openRightSidebar() {
    this.isRightSidebarOpen = true;
  }

  openDiagnosticsPanel() {
    this.isRightSidebarOpen = true;
    this.rightSidebarTab = 'issues';
  }

  openNodePanel() {
    this.isRightSidebarOpen = true;
    this.rightSidebarTab = 'node';
  }

  openConfigPanel() {
    this.isRightSidebarOpen = true;
    this.rightSidebarTab = 'project';
  }

  /** openConfirmModal stages a confirm dialog with the given title,
   * message, and on-confirm callback. The callback runs when the user
   * presses Confirm; Cancel just closes the modal without calling it.
   */
  openConfirmModal(title: string, message: string, onConfirm: () => void) {
    this.confirmProps = { title, message, onConfirm };
    this.modalKind = 'confirm';
  }

  /** openInputModal stages a text-input prompt. onConfirm receives the
   * value when the user presses Confirm (or hits Enter).
   */
  openInputModal(title: string, placeholder: string, onConfirm: (value: string) => void) {
    this.inputProps = { title, placeholder, onConfirm };
    this.modalKind = 'input';
  }

  /** openSettingsModal opens the global settings dialog overlay. */
  openSettingsModal() {
    this.modalKind = 'settings';
  }

  /** closeModal dismisses whichever modal is currently active. */
  closeModal() {
    this.modalKind = null;
  }

  showToast(kind: 'error' | 'info', message: string) {
    const id = Date.now();
    this.toast = { id, kind, message };
    setTimeout(() => {
      if (this.toast?.id === id) this.toast = null;
    }, 4000);
  }
}

export const applicationStore = new ApplicationStore();
