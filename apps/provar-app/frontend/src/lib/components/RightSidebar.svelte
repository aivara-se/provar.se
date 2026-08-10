<script lang="ts">
  import { editorStore } from '../stores/editor-store.svelte';
  import { uiStore } from '../stores/ui-store.svelte';
  import NodeSidePanel from './NodeSidePanel.svelte';
  import ProjectConfigPanel from './ProjectConfigPanel.svelte';

  type PanelKind = 'config' | 'node';

  // Active panel: a selected node wins, otherwise the config panel.
  let active = $derived<PanelKind>(
    editorStore.selectedNodeId ? 'node' : 'config',
  );
</script>

{#if uiStore.isRightSidebarOpen && active}
  <aside
    class="absolute top-0 right-0 bottom-0 z-20 flex w-[400px] flex-col border-l border-zinc-800 bg-[#161b22]/50 pt-[4px] backdrop-blur-md"
  >
    {#if active === 'config'}
      <ProjectConfigPanel />
    {:else}
      <NodeSidePanel />
    {/if}
  </aside>
{/if}