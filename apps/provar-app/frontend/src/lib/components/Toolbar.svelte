<script lang="ts">
  import { PanelLeft, Settings, Play, Hammer, Square, RotateCw } from 'lucide-svelte';
  import { editorStore } from '../stores/editor-store.svelte';
  import { uiStore } from '../stores/ui-store.svelte';

  let fileName = $derived(
    editorStore.selectedFilePath
      ? editorStore.selectedFilePath.split('/').pop()
      : null,
  );

  let hasFile = $derived(editorStore.selectedFilePath !== null);

  // Run and Compile are mutually exclusive on the toolbar — the engine
  // allows concurrent jobs across the bind, but the editor's per-file
  // model wants one action at a time. The bind layers also wouldn't
  // be safe; we let the engine.Rules do the right thing and just
  // gate the buttons here.
  let busy = $derived(editorStore.isRunning || editorStore.isCompiling);
</script>

<div
  class="absolute top-0 right-0 left-0 z-30 flex h-12 items-center gap-3 pl-[88px] pr-3 pt-[4px] text-xs text-zinc-400"
>
  <button
    type="button"
    class="rounded p-1.5 text-zinc-400 transition-colors hover:bg-zinc-800 hover:text-zinc-200"
    title={uiStore.isSidebarOpen ? 'Hide Sidebar' : 'Show Sidebar'}
    onclick={() => uiStore.toggleSidebar()}
  >
    <PanelLeft class="h-4 w-4" />
  </button>

  <div class="flex-1"></div>

  {#if fileName}
    <span class="font-mono text-zinc-400">{fileName}</span>
  {/if}

  <div class="flex-1"></div>

  {#if hasFile}
    {#if editorStore.isCompiling || Object.keys(editorStore.compileStates).length > 0}
      <div
        class="flex items-center gap-1.5 rounded border border-zinc-800/80 bg-[#161b22]/80 px-2.5 py-1 text-xs text-zinc-300 backdrop-blur-sm"
        title="Compiling…"
      >
        <div
          class="h-3 w-3 animate-spin rounded-full border border-zinc-500 border-t-amber-500"
        ></div>
        <span>Compiling…</span>
      </div>
    {:else}
      <button
        type="button"
        class="flex items-center gap-1 rounded p-1.5 text-zinc-400 transition-colors hover:bg-zinc-800 hover:text-zinc-200"
        title="Compile test"
        disabled={busy}
        onclick={() => void editorStore.compileCurrent()}
      >
        <Hammer class="h-3.5 w-3.5" />
      </button>
    {/if}

    {#if editorStore.isRunning || editorStore.isCompiling}
      <button
        type="button"
        class="flex items-center gap-1 rounded border border-red-900/60 bg-[#161b22]/80 px-2.5 py-1 text-xs font-medium text-red-400 backdrop-blur-sm transition-colors hover:bg-red-950/40"
        title={editorStore.isRunning ? 'Stop running test' : 'Stop compiling'}
        onclick={() => {
          if (editorStore.isRunning) void editorStore.stopRun();
          else void editorStore.stopCompile();
        }}
      >
        <Square class="h-3 w-3 fill-current" />
        <span>Stop</span>
      </button>
    {:else}
      <button
        type="button"
        class="flex items-center gap-1 rounded p-1.5 text-zinc-400 transition-colors hover:bg-zinc-800 hover:text-green-400"
        title="Run test"
        disabled={busy}
        onclick={() => void editorStore.runCurrent()}
      >
        <Play class="h-3.5 w-3.5 fill-current" />
      </button>
    {/if}

    <button
      type="button"
      class="flex items-center gap-1 rounded p-1.5 text-zinc-500 transition-colors hover:bg-zinc-800 hover:text-zinc-300"
      title="Project Settings"
      onclick={() => uiStore.openRightSidebar()}
    >
      <Settings class="h-3.5 w-3.5" />
    </button>
  {/if}
</div>
