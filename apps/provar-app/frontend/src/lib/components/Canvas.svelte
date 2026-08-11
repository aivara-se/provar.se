<script lang="ts">
  import { editorStore } from '../stores/editor-store.svelte';
  import { applicationStore } from '../stores/application-store.svelte';
  import Canvas from '../canvas/Canvas.svelte';
</script>

{#if editorStore.currentFile}
  <Canvas />
{:else}
  <!-- svelte-ignore a11y_click_events_have_key_events -->
  <!-- svelte-ignore a11y_no_static_element_interactions -->
  <div
    class="absolute inset-0 flex items-center justify-center pt-[48px] text-sm text-zinc-500"
    onclick={() => applicationStore.closeSidebars()}
  >
    <p>Select a test to begin</p>
  </div>
{/if}