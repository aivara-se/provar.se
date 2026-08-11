<script lang="ts">
  import { editorStore } from '../../stores/editor-store.svelte';
  import { executionStore } from '../../stores/execution-store.svelte';
  import { applicationStore, type RightSidebarTab } from '../../stores/application-store.svelte';
  import NodeSidePanel from './NodeSidePanel.svelte';
  import ProjectConfigPanel from './ProjectConfigPanel.svelte';
  import DiagnosticsSidePanel from './DiagnosticsSidePanel.svelte';
  import ConsolePanel from './ConsolePanel.svelte';

  let issueCount = $derived(
    editorStore.diagnostics.errors.length + editorStore.diagnostics.warnings.length,
  );
  let logCount = $derived(executionStore.logs.length);

  function setTab(tab: RightSidebarTab) {
    applicationStore.rightSidebarTab = tab;
  }
</script>

{#if applicationStore.isRightSidebarOpen}
  <aside
    class="absolute top-0 right-0 bottom-0 z-30 flex w-[400px] flex-col border-l border-zinc-800 bg-[#161b22]/95 backdrop-blur-md text-zinc-300 shadow-2xl"
  >
    <!-- Tab Navigation Header Bar (clean text-only tabs, no dark background, no borders, no icons) -->
    <div class="flex h-11 shrink-0 items-center px-4 pt-2 select-none">
      <nav class="flex items-center gap-4 text-xs">
        <!-- Project Tab -->
        <button
          type="button"
          onclick={() => setTab('project')}
          class="cursor-pointer transition-colors focus:outline-none {applicationStore.rightSidebarTab === 'project' ? 'font-semibold text-zinc-100' : 'text-zinc-500 hover:text-zinc-300'}"
        >
          Project
        </button>

        <!-- Issues Tab -->
        <button
          type="button"
          onclick={() => setTab('issues')}
          class="flex items-center gap-1 cursor-pointer transition-colors focus:outline-none {applicationStore.rightSidebarTab === 'issues' ? 'font-semibold text-zinc-100' : 'text-zinc-500 hover:text-zinc-300'}"
        >
          <span>Issues</span>
          {#if issueCount > 0}
            <span class="text-[10px] text-amber-400 font-semibold">({issueCount})</span>
          {/if}
        </button>

        <!-- Actions Tab -->
        <button
          type="button"
          onclick={() => setTab('node')}
          class="cursor-pointer transition-colors focus:outline-none {applicationStore.rightSidebarTab === 'node' ? 'font-semibold text-zinc-100' : 'text-zinc-500 hover:text-zinc-300'}"
        >
          Actions
        </button>

        <!-- Console Tab -->
        <button
          type="button"
          onclick={() => setTab('console')}
          class="flex items-center gap-1 cursor-pointer transition-colors focus:outline-none {applicationStore.rightSidebarTab === 'console' ? 'font-semibold text-zinc-100' : 'text-zinc-500 hover:text-zinc-300'}"
        >
          <span>Console</span>
          {#if logCount > 0}
            <span class="text-[10px] text-blue-400 font-semibold">({logCount})</span>
          {/if}
        </button>
      </nav>
    </div>

    <!-- Active Panel Body -->
    <div class="flex-1 overflow-hidden">
      {#if applicationStore.rightSidebarTab === 'issues'}
        <DiagnosticsSidePanel />
      {:else if applicationStore.rightSidebarTab === 'node'}
        <NodeSidePanel />
      {:else if applicationStore.rightSidebarTab === 'console'}
        <ConsolePanel />
      {:else}
        <ProjectConfigPanel />
      {/if}
    </div>
  </aside>
{/if}