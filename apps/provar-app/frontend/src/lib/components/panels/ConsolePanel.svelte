<script lang="ts">
  import {
    Terminal,
    Trash2,
    Search,
    ChevronRight,
    ChevronDown,
    ArrowDownToLine,
  } from 'lucide-svelte';
  import { executionStore } from '../../stores/execution-store.svelte';
  import { applicationStore } from '../../stores/application-store.svelte';
  import { editorStore } from '../../stores/editor-store.svelte';

  type FilterLevel = 'all' | 'info' | 'step' | 'success' | 'warn' | 'error';

  let filter = $state<FilterLevel>('all');
  let searchQuery = $state('');
  let autoScroll = $state(true);
  let expandedDetails = $state<Record<string, boolean>>({});
  let terminalContainerEl = $state<HTMLDivElement>();

  let filteredLogs = $derived.by(() => {
    let result = executionStore.logs;
    if (filter !== 'all') {
      result = result.filter((log) => log.level === filter);
    }
    if (searchQuery.trim()) {
      const q = searchQuery.toLowerCase();
      result = result.filter(
        (log) =>
          log.message.toLowerCase().includes(q) ||
          (log.actionId && log.actionId.toLowerCase().includes(q)) ||
          (log.details && log.details.toLowerCase().includes(q)) ||
          (log.category && log.category.toLowerCase().includes(q)),
      );
    }
    return result;
  });

  let counts = $derived.by(() => {
    const logs = executionStore.logs;
    return {
      all: logs.length,
      info: logs.filter((l) => l.level === 'info').length,
      step: logs.filter((l) => l.level === 'step').length,
      success: logs.filter((l) => l.level === 'success').length,
      warn: logs.filter((l) => l.level === 'warn').length,
      error: logs.filter((l) => l.level === 'error').length,
    };
  });

  let statusBadge = $derived.by(() => {
    if (executionStore.isRunning)
      return { text: 'Running', color: 'text-emerald-400 bg-emerald-500/10 border-emerald-500/30 animate-pulse' };
    if (executionStore.isCompiling)
      return { text: 'Compiling', color: 'text-blue-400 bg-blue-500/10 border-blue-500/30 animate-pulse' };
    if (counts.error > 0) return { text: 'Failed', color: 'text-red-400 bg-red-500/10 border-red-500/30' };
    if (counts.success > 0) return { text: 'Success', color: 'text-emerald-400 bg-emerald-500/10 border-emerald-500/30' };
    return { text: 'Idle', color: 'text-zinc-500 bg-zinc-800/50 border-zinc-700/50' };
  });

  $effect(() => {
    if (autoScroll && filteredLogs.length > 0 && terminalContainerEl) {
      setTimeout(() => {
        if (terminalContainerEl) {
          terminalContainerEl.scrollTop = terminalContainerEl.scrollHeight;
        }
      }, 10);
    }
  });

  function toggleDetails(logId: string) {
    expandedDetails = { ...expandedDetails, [logId]: !expandedDetails[logId] };
  }

  function handleNodeClick(actionId?: string) {
    if (!actionId) return;
    editorStore.selectedNodeId = actionId;
    applicationStore.openNodePanel();
  }
</script>

<div class="flex h-full flex-col text-zinc-300">
  <!-- Controls bar -->
  <div class="shrink-0 flex flex-col gap-2 border-b border-zinc-800/50 p-3 select-none">
    <div class="flex items-center justify-between">
      <div class="flex items-center gap-2">
        <span class="rounded-full border px-2 py-0.5 text-[10px] font-semibold tracking-wide {statusBadge.color}">
          {statusBadge.text}
        </span>
        <span class="text-xs text-zinc-500">
          {counts.all} {counts.all === 1 ? 'entry' : 'entries'}
        </span>
      </div>

      <div class="flex items-center gap-1">
        <!-- Auto Scroll Toggle -->
        <button
          type="button"
          onclick={() => (autoScroll = !autoScroll)}
          class="flex h-6 w-6 cursor-pointer items-center justify-center rounded transition-colors {autoScroll ? 'text-blue-400' : 'text-zinc-500 hover:text-zinc-300'}"
          title={autoScroll ? 'Auto-scroll enabled' : 'Auto-scroll disabled'}
        >
          <ArrowDownToLine class="h-3.5 w-3.5" />
        </button>

        <!-- Clear Logs -->
        <button
          type="button"
          onclick={() => executionStore.clearLogs()}
          class="flex h-6 w-6 cursor-pointer items-center justify-center rounded text-zinc-500 transition-colors hover:text-zinc-300"
          title="Clear console output"
        >
          <Trash2 class="h-3.5 w-3.5" />
        </button>
      </div>
    </div>

    <!-- Search Input -->
    <div class="relative flex items-center">
      <Search class="absolute left-2.5 h-3 w-3 text-zinc-500" />
      <input
        type="text"
        bind:value={searchQuery}
        placeholder="Filter logs..."
        class="h-7 w-full rounded border border-zinc-800/60 bg-zinc-900/40 pl-8 pr-2.5 text-xs text-zinc-200 placeholder-zinc-600 focus:border-zinc-700 focus:outline-none"
      />
    </div>

    <!-- Filter Tabs -->
    <div class="flex flex-wrap items-center gap-1">
      <button
        type="button"
        onclick={() => (filter = 'all')}
        class="rounded px-2 py-0.5 text-[11px] transition-colors {filter === 'all' ? 'text-zinc-100 font-medium' : 'text-zinc-500 hover:text-zinc-300'}"
      >
        All ({counts.all})
      </button>
      <button
        type="button"
        onclick={() => (filter = 'step')}
        class="rounded px-2 py-0.5 text-[11px] transition-colors {filter === 'step' ? 'text-purple-300 font-medium' : 'text-zinc-500 hover:text-zinc-300'}"
      >
        Steps ({counts.step})
      </button>
      <button
        type="button"
        onclick={() => (filter = 'info')}
        class="rounded px-2 py-0.5 text-[11px] transition-colors {filter === 'info' ? 'text-blue-300 font-medium' : 'text-zinc-500 hover:text-zinc-300'}"
      >
        Info ({counts.info})
      </button>
      <button
        type="button"
        onclick={() => (filter = 'success')}
        class="rounded px-2 py-0.5 text-[11px] transition-colors {filter === 'success' ? 'text-emerald-300 font-medium' : 'text-zinc-500 hover:text-zinc-300'}"
      >
        Success ({counts.success})
      </button>
      {#if counts.error > 0}
        <button
          type="button"
          onclick={() => (filter = 'error')}
          class="rounded px-2 py-0.5 text-[11px] transition-colors {filter === 'error' ? 'text-red-300 font-medium' : 'text-red-400/70 hover:text-red-300'}"
        >
          Errors ({counts.error})
        </button>
      {/if}
    </div>
  </div>

  <!-- Log Messages -->
  <div
    bind:this={terminalContainerEl}
    class="flex-1 overflow-y-auto p-3 font-mono text-[11px] leading-relaxed select-text"
  >
    {#if filteredLogs.length === 0}
      <div class="flex h-full flex-col items-center justify-center gap-2 text-center text-zinc-600 select-none py-12">
        <Terminal class="h-7 w-7 text-zinc-700" />
        <p class="text-xs text-zinc-500">No console output available.</p>
        <p class="text-[11px] text-zinc-600">Run or compile a test to view logs.</p>
      </div>
    {:else}
      <div class="flex flex-col gap-1">
        {#each filteredLogs as log (log.id)}
          <div class="group flex flex-col rounded p-1.5 transition-colors hover:bg-zinc-800/30">
            <div class="flex flex-wrap items-baseline gap-1.5">
              <!-- Timestamp -->
              <span class="shrink-0 text-zinc-600 text-[10px]">{log.timestamp}</span>

              <!-- Level Badge -->
              {#if log.level === 'info'}
                <span class="shrink-0 rounded px-1 text-[9px] font-semibold text-blue-400">INFO</span>
              {:else if log.level === 'step'}
                <span class="shrink-0 rounded px-1 text-[9px] font-semibold text-purple-400">STEP</span>
              {:else if log.level === 'success'}
                <span class="shrink-0 rounded px-1 text-[9px] font-semibold text-emerald-400">OK</span>
              {:else if log.level === 'warn'}
                <span class="shrink-0 rounded px-1 text-[9px] font-semibold text-amber-400">WARN</span>
              {:else if log.level === 'error'}
                <span class="shrink-0 rounded px-1 text-[9px] font-semibold text-red-400">FAIL</span>
              {/if}

              <!-- Category Badge -->
              {#if log.category}
                <span class="shrink-0 text-[9px] uppercase font-bold text-zinc-600">
                  [{log.category}]
                </span>
              {/if}

              <!-- Action Node Badge -->
              {#if log.actionId}
                <button
                  type="button"
                  onclick={() => handleNodeClick(log.actionId)}
                  class="shrink-0 cursor-pointer rounded px-1 text-[9px] font-medium text-blue-400 hover:underline"
                  title="Click to select node in editor"
                >
                  node:{log.actionId}
                </button>
              {/if}

              <!-- Details Toggle -->
              {#if log.details}
                <button
                  type="button"
                  onclick={() => toggleDetails(log.id)}
                  class="ml-auto flex items-center gap-0.5 shrink-0 text-[10px] text-zinc-600 hover:text-zinc-400 transition-colors"
                >
                  {#if expandedDetails[log.id]}
                    <ChevronDown class="h-3 w-3" />
                  {:else}
                    <ChevronRight class="h-3 w-3" />
                  {/if}
                </button>
              {/if}
            </div>

            <!-- Message Body -->
            <div class="mt-0.5 text-zinc-300 break-words leading-normal {log.level === 'error' ? 'text-red-300' : ''}">
              {log.message}
            </div>

            <!-- Expanded Details -->
            {#if log.details && expandedDetails[log.id]}
              <div class="mt-1.5 overflow-x-auto rounded border border-zinc-800/60 bg-zinc-900/40 p-2 text-[10px] text-zinc-400 whitespace-pre-wrap">
                {log.details}
              </div>
            {/if}
          </div>
        {/each}
      </div>
    {/if}
  </div>
</div>
