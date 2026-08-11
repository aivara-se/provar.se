<script lang="ts">
  import { applicationStore } from '../../stores/application-store.svelte';
  import SettingsModal from './SettingsModal.svelte';
  import ConfigModal from './ConfigModal.svelte';
  import ConfirmModal from './ConfirmModal.svelte';
  import InputModal from './InputModal.svelte';
  import DoctorModal from './DoctorModal.svelte';

  // AppModals hosts every overlay rendered above the editor. Settings and
  // Config are permanent (visible whenever a project is open); Confirm
  // and Input are dispatched by applicationStore.modalKind — callers stage the
  // payload with applicationStore.openConfirmModal / openInputModal.
</script>

<SettingsModal />
<ConfigModal />
<ConfirmModal />
<InputModal />
<DoctorModal />

<!--
  Toast is rendered at the App level so it floats above every panel.
  The store's timeout (4s) auto-clears; new toasts replace the current
  one and reset the timer.
-->
{#if applicationStore.toast}
  <div
    class="pointer-events-none fixed right-4 bottom-4 z-[300] flex max-w-sm items-start gap-2 rounded-lg border px-4 py-3 text-sm shadow-2xl backdrop-blur-md transition-all {applicationStore
      .toast.kind === 'error'
      ? 'border-red-500/40 bg-red-950/80 text-red-200'
      : 'border-zinc-700 bg-zinc-900/80 text-zinc-200'}"
    role="status"
    aria-live="polite"
    data-testid="app-toast"
  >
    <span class="font-mono text-[10px] tracking-wider uppercase opacity-70">
      {applicationStore.toast.kind}
    </span>
    <span>{applicationStore.toast.message}</span>
  </div>
{/if}
