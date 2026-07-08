<script lang="ts">
  import Modal from './Modal.svelte';
  import { uiStore } from '../stores/ui-store.svelte';

  // Confirm is a thin wrapper over the shared Modal. It reads all of its
  // payload off the uiStore — no props — so callers stage the dialog with
  // openConfirmModal(title, message, onConfirm) and forget about it.
  //
  // Cancel is wired to close the modal without calling the user's
  // callback. Backdrop-click also closes (Modal fires onClose).
  function handleCancel() {
    uiStore.closeModal();
  }

  function handleConfirm() {
    const props = uiStore.confirmProps;
    uiStore.closeModal();
    props.onConfirm();
  }
</script>

<Modal
  show={uiStore.modalKind === 'confirm'}
  title={uiStore.confirmProps.title}
  primaryLabel="Confirm"
  onPrimary={handleConfirm}
  onClose={handleCancel}
  primaryStyle="danger"
>
  <p class="text-sm leading-relaxed text-zinc-400">{uiStore.confirmProps.message}</p>
</Modal>
