<script lang="ts">
  import { tick } from 'svelte';
  import Modal from './Modal.svelte';
  import { uiStore } from '../stores/ui-store.svelte';

  let value = $state('');
  let input: HTMLInputElement | undefined = $state();

  // Focus the input when the prompt opens, and clear any leftover text
  // from a previous prompt when it closes so the user doesn't see
  // stale state flash by.
  $effect(() => {
    if (uiStore.modalKind === 'input') {
      tick().then(() => input?.focus());
    } else {
      value = '';
    }
  });

  function handleCancel() {
    uiStore.closeModal();
  }

  function handleConfirm() {
    const next = value;
    const props = uiStore.inputProps;
    uiStore.closeModal();
    props.onConfirm(next);
  }
</script>

<Modal
  show={uiStore.modalKind === 'input'}
  title={uiStore.inputProps.title}
  primaryLabel="Confirm"
  onPrimary={handleConfirm}
  onClose={handleCancel}
>
  <input
    type="text"
    bind:value
    bind:this={input}
    placeholder={uiStore.inputProps.placeholder}
    class="w-full rounded-lg border border-zinc-700/50 bg-[#21262d] p-2.5 text-sm text-zinc-200 placeholder-zinc-500 focus:border-indigo-500 focus:ring-1 focus:ring-indigo-500 focus:outline-none"
    onkeydown={(e) => {
      if (e.key === 'Enter') handleConfirm();
    }}
  />
</Modal>
