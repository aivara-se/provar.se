<script lang="ts">
  import Modal from './Modal.svelte';
  import { applicationStore } from '../stores/application-store.svelte';

  // Confirm is a thin wrapper over the shared Modal. It reads all of its
  // payload off the applicationStore — no props — so callers stage the dialog with
  // openConfirmModal(title, message, onConfirm) and forget about it.
  //
  // Cancel is wired to close the modal without calling the user's
  // callback. Backdrop-click also closes (Modal fires onClose).
  function handleCancel() {
    applicationStore.closeModal();
  }

  function handleConfirm() {
    const props = applicationStore.confirmProps;
    applicationStore.closeModal();
    props.onConfirm();
  }
</script>

<Modal
  show={applicationStore.modalKind === 'confirm'}
  title={applicationStore.confirmProps.title}
  primaryLabel="Confirm"
  onPrimary={handleConfirm}
  onClose={handleCancel}
  primaryStyle="danger"
>
  <p class="text-sm leading-relaxed text-zinc-400">{applicationStore.confirmProps.message}</p>
</Modal>
