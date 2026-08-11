<script lang="ts">
  import { editorStore } from '../../stores/editor-store.svelte';
  import { applicationStore } from '../../stores/application-store.svelte';
  import NodeSidePanel from './NodeSidePanel.svelte';
  import ProjectConfigPanel from './ProjectConfigPanel.svelte';
  import DiagnosticsSidePanel from './DiagnosticsSidePanel.svelte';

  type PanelKind = 'config' | 'node' | 'diagnostics';

  let active = $derived.by<PanelKind>(() => {
    if (applicationStore.rightSidebarTab === 'diagnostics') {
      return 'diagnostics';
    }
    return editorStore.selectedNodeId ? 'node' : 'config';
  });
</script>

{#if applicationStore.isRightSidebarOpen && active}
  <aside
    class="absolute top-0 right-0 bottom-0 z-20 flex w-[400px] flex-col border-l border-zinc-800 bg-[#161b22]/50 pt-[4px] backdrop-blur-md"
  >
    {#if active === 'diagnostics'}
      <DiagnosticsSidePanel />
    {:else if active === 'node'}
      <NodeSidePanel />
    {:else}
      <ProjectConfigPanel />
    {/if}
  </aside>
{/if}