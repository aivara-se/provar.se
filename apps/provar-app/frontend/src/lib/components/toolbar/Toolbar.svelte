<script lang="ts">
  import {
    File as FileIcon,
    Play,
    Hammer,
    Square,
    ChevronDown,
    X,
    AlertTriangle,
  } from 'lucide-svelte';
  import { editorStore } from '../../stores/editor-store.svelte';
  import { executionStore } from '../../stores/execution-store.svelte';
  import { applicationStore } from '../../stores/application-store.svelte';

  let runMenuOpen = $state(false);
  let toolbarEl = $state<HTMLDivElement>();

  let fileName = $derived.by(() => {
    if (!editorStore.selectedFilePath) return null;
    return editorStore.selectedFilePath.replace(/^\.provar\/tests\//, '');
  });

  let busy = $derived(executionStore.isRunning || executionStore.isCompiling);
  let hasDiagnosticsError = $derived(!editorStore.diagnostics.isValid);

  function handleDocumentClick(e: MouseEvent) {
    if (!runMenuOpen) return;
    const target = e.target as Node | null;
    if (target && !toolbarEl?.contains(target)) {
      runMenuOpen = false;
    }
  }

  $effect(() => {
    document.addEventListener('click', handleDocumentClick);
    return () => document.removeEventListener('click', handleDocumentClick);
  });
</script>

{#if editorStore.selectedFilePath}
  <div
    bind:this={toolbarEl}
    class="pointer-events-auto absolute top-[8px] left-1/2 z-50 flex h-[26px] -translate-x-1/2 items-center rounded-full border border-zinc-800/80 bg-[#161b22]/80 shadow-sm backdrop-blur-sm divide-x divide-zinc-800/80 text-xs select-none"
  >
    <div
      class="flex h-full items-center gap-1.5 rounded-l-full px-3 py-1 text-xs font-medium text-zinc-300"
    >
      <FileIcon class="h-3.5 w-3.5 text-blue-400" />
      <span class="tracking-wide">{fileName}</span>
      {#if !editorStore.diagnostics.isValid || editorStore.diagnostics.warnings.length > 0}
        {@const issueCount = editorStore.diagnostics.errors.length + editorStore.diagnostics.warnings.length}
        <button
          type="button"
          onclick={() => applicationStore.openDiagnosticsPanel()}
          class="flex items-center gap-1 text-xs font-medium text-amber-400 transition-colors hover:text-amber-300"
          title="Graph has validation issues. Click to open Diagnostics Panel."
        >
          <AlertTriangle class="h-3.5 w-3.5 shrink-0 text-amber-400" />
          <span>{issueCount}</span>
        </button>
      {/if}
    </div>

    {#if executionStore.isCompiling}
      <div class="flex h-full items-center gap-1.5 rounded-r-full px-3 py-1 text-xs text-zinc-400">
        <div
          class="h-3 w-3 animate-spin rounded-full border border-zinc-500 border-t-blue-500"
        ></div>
        <span>Compiling...</span>
      </div>
    {:else if executionStore.isRunning}
      <button
        type="button"
        onclick={() => void editorStore.stopRun()}
        class="flex h-full cursor-pointer items-center gap-1.5 rounded-r-full px-3 py-1 text-xs font-medium text-red-400 transition-colors duration-200 hover:bg-red-950/40 hover:text-red-300 focus:outline-none"
        title="Stop running test"
      >
        <Square class="h-2.5 w-2.5 fill-current" />
        <span>Stop</span>
      </button>
    {:else}
      <div class="relative flex h-full items-center divide-x divide-zinc-800/80">
        {#if editorStore.needsCompile}
          <button
            type="button"
            onclick={() => void editorStore.compileCurrent()}
            disabled={busy}
            class="flex h-full cursor-pointer items-center gap-1.5 px-3 py-1 text-xs font-medium text-zinc-400 transition-colors duration-200 hover:bg-[#21262d]/90 hover:text-blue-400 focus:outline-none disabled:opacity-50"
            title="Compile test"
          >
            <Hammer class="h-3 w-3" />
            <span>Compile</span>
          </button>
        {:else}
          <button
            type="button"
            onclick={() => void editorStore.runCurrent()}
            disabled={busy || hasDiagnosticsError}
            class="flex h-full cursor-pointer items-center gap-1.5 px-3 py-1 text-xs font-medium text-zinc-400 transition-colors duration-200 hover:bg-[#21262d]/90 hover:text-green-400 focus:outline-none disabled:opacity-50"
            title={hasDiagnosticsError ? "Fix graph diagnostics before running" : "Run test"}
          >
            <Play class="h-2.5 w-2.5 fill-current" />
            <span>Run</span>
          </button>
        {/if}

        <button
          type="button"
          onclick={(e) => {
            e.stopPropagation();
            runMenuOpen = !runMenuOpen;
          }}
          class="flex h-full cursor-pointer items-center rounded-r-full px-1.5 text-zinc-500 transition-colors duration-200 hover:bg-[#21262d]/90 hover:text-zinc-300 focus:outline-none"
          title="Run options"
        >
          <ChevronDown class="h-3 w-3" />
        </button>

        {#if runMenuOpen}
          <div
            class="absolute top-[30px] right-0 z-50 min-w-[140px] overflow-hidden rounded-lg border border-zinc-800 bg-[#161b22] shadow-xl"
          >
            <button
              type="button"
              disabled={busy || hasDiagnosticsError}
              onclick={() => {
                runMenuOpen = false;
                void editorStore.runCurrent();
              }}
              class="flex w-full items-center gap-2 px-3 py-2 text-left text-xs text-zinc-300 transition-colors hover:bg-zinc-800/60 disabled:opacity-50"
            >
              <Play class="h-3 w-3 shrink-0 text-zinc-400 fill-current" />
              <span>Run test</span>
            </button>
            <button
              type="button"
              disabled={busy}
              onclick={() => {
                runMenuOpen = false;
                void editorStore.compileCurrent();
              }}
              class="flex w-full items-center gap-2 px-3 py-2 text-left text-xs text-zinc-400 transition-colors hover:bg-zinc-800/60 hover:text-zinc-300 disabled:opacity-50"
            >
              <Hammer class="h-3 w-3 shrink-0" />
              <span>Compile test</span>
            </button>
            <button
              type="button"
              disabled={busy}
              onclick={() => {
                runMenuOpen = false;
                executionStore.clearStates();
              }}
              class="flex w-full items-center gap-2 px-3 py-2 text-left text-xs text-zinc-400 transition-colors hover:bg-zinc-800/60 hover:text-zinc-300 disabled:opacity-50"
            >
              <X class="h-3 w-3 shrink-0" />
              <span>Clear status</span>
            </button>
          </div>
        {/if}
      </div>
    {/if}
  </div>
{/if}
