<script lang="ts">
  import {
    Terminal,
    X,
    Trash2,
    Search,
    ChevronRight,
    ChevronDown,
    Maximize2,
    Minimize2,
    ArrowDownToLine,
  } from 'lucide-svelte';
  import { executionStore, type LogEntry } from '../../stores/execution-store.svelte';
  import { applicationStore } from '../../stores/application-store.svelte';
  import { editorStore } from '../../stores/editor-store.svelte';

  type FilterLevel = 'all' | 'info' | 'step' | 'success' | 'warn' | 'error';

  let filter = $state<FilterLevel>('all');
  let searchQuery = $state('');
  let autoScroll = $state(true);
  let isMaximized = $state(false);
  let expandedDetails = $state<Record<string, boolean>>({});
  let drawerHeight = $state(260);
  let isDragging = $state(false);
  let startY = 0;
  let startHeight = 0;

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
    if (executionStore.isRunning) return { text: 'Running', color: 'text-emerald-400 bg-emerald-500/10 border-emerald-500/30 animate-pulse' };
    if (executionStore.isCompiling) return { text: 'Compiling', color: 'text-blue-400 bg-blue-500/10 border-blue-500/30 animate-pulse' };
    if (counts.error > 0) return { text: 'Failed', color: 'text-red-400 bg-red-500/10 border-red-500/30' };
    if (counts.success > 0) return { text: 'Success', color: 'text-emerald-400 bg-emerald-500/10 border-emerald-500/30' };
    return { text: 'Idle', color: 'text-zinc-500 bg-zinc-800/50 border-zinc-700/50' };
  });

  $effect(() => {
    // Auto-scroll terminal when new logs arrive if autoScroll is enabled
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

  function startResize(e: MouseEvent) {
    isDragging = true;
    startY = e.clientY;
    startHeight = drawerHeight;
    document.addEventListener('mousemove', handleResize);
    document.addEventListener('mouseup', stopResize);
  }

  function handleResize(e: MouseEvent) {
    if (!isDragging) return;
    const deltaY = startY - e.clientY;
    const newHeight = Math.min(Math.max(startHeight + deltaY, 140), window.innerHeight * 0.75);
    drawerHeight = newHeight;
  }

  function stopResize() {
    isDragging = false;
    document.removeEventListener('mousemove', handleResize);
    document.removeEventListener('mouseup', stopResize);
  }
</script>

{#if applicationStore.isConsoleOpen}
  <div
    class="absolute right-0 bottom-0 left-0 z-40 flex flex-col border-t border-zinc-800 bg-[#0d1117] text-zinc-300 shadow-2xl transition-all duration-150"
    style="height: {isMaximized ? 'calc(100vh - 40px)' : `${drawerHeight}px`}; left: {applicationStore.isSidebarOpen ? '256px' : '0px'};"
  >
    <!-- Top Drag Handle -->
    {#if !isMaximized}
      <!-- svelte-ignore a11y_no_static_element_interactions -->
      <div
        onmousedown={startResize}
        class="h-1.5 w-full cursor-row-resize bg-transparent hover:bg-blue-500/40 transition-colors"
        title="Drag to resize drawer height"
      ></div>
    {/if}

    <!-- Header Toolbar -->
    <div
      class="flex h-9 shrink-0 items-center justify-between border-b border-zinc-800/80 bg-[#161b22] px-3 select-none text-xs"
    >
      <div class="flex items-center gap-3">
        <div class="flex items-center gap-1.5 font-medium text-zinc-200">
          <Terminal class="h-3.5 w-3.5 text-blue-400" />
          <span>Output Console</span>
        </div>

        <span class="rounded-full border px-2 py-0.5 text-[10px] font-semibold tracking-wide {statusBadge.color}">
          {statusBadge.text}
        </span>

        <!-- Filter Tabs -->
        <div class="ml-2 flex items-center gap-1">
          <button
            type="button"
            onclick={() => (filter = 'all')}
            class="rounded px-2 py-0.5 text-[11px] transition-colors {filter === 'all' ? 'bg-zinc-800 text-zinc-100 font-medium' : 'text-zinc-400 hover:text-zinc-200'}"
          >
            All ({counts.all})
          </button>
          <button
            type="button"
            onclick={() => (filter = 'step')}
            class="rounded px-2 py-0.5 text-[11px] transition-colors {filter === 'step' ? 'bg-purple-950/60 text-purple-300 font-medium' : 'text-zinc-400 hover:text-zinc-200'}"
          >
            Steps ({counts.step})
          </button>
          <button
            type="button"
            onclick={() => (filter = 'info')}
            class="rounded px-2 py-0.5 text-[11px] transition-colors {filter === 'info' ? 'bg-blue-950/60 text-blue-300 font-medium' : 'text-zinc-400 hover:text-zinc-200'}"
          >
            Info ({counts.info})
          </button>
          <button
            type="button"
            onclick={() => (filter = 'success')}
            class="rounded px-2 py-0.5 text-[11px] transition-colors {filter === 'success' ? 'bg-emerald-950/60 text-emerald-300 font-medium' : 'text-zinc-400 hover:text-zinc-200'}"
          >
            Success ({counts.success})
          </button>
          {#if counts.error > 0}
            <button
              type="button"
              onclick={() => (filter = 'error')}
              class="rounded px-2 py-0.5 text-[11px] transition-colors {filter === 'error' ? 'bg-red-950/60 text-red-300 font-medium' : 'text-red-400/80 hover:text-red-300'}"
            >
              Errors ({counts.error})
            </button>
          {/if}
        </div>
      </div>

      <!-- Right Controls -->
      <div class="flex items-center gap-2">
        <!-- Search Input -->
        <div class="relative flex items-center">
          <Search class="absolute left-2 h-3 w-3 text-zinc-500" />
          <input
            type="text"
            bind:value={searchQuery}
            placeholder="Filter logs..."
            class="h-6 w-36 rounded border border-zinc-800 bg-[#0d1117] pl-7 pr-2 text-[11px] text-zinc-200 placeholder-zinc-500 focus:border-blue-500/50 focus:outline-none"
          />
        </div>

        <!-- Auto Scroll Toggle -->
        <button
          type="button"
          onclick={() => (autoScroll = !autoScroll)}
          class="flex h-6 w-6 cursor-pointer items-center justify-center rounded transition-colors {autoScroll ? 'text-blue-400 bg-blue-950/40' : 'text-zinc-500 hover:text-zinc-300'}"
          title={autoScroll ? 'Auto-scroll enabled' : 'Auto-scroll disabled'}
        >
          <ArrowDownToLine class="h-3.5 w-3.5" />
        </button>

        <!-- Clear Button -->
        <button
          type="button"
          onclick={() => executionStore.clearLogs()}
          class="flex h-6 w-6 cursor-pointer items-center justify-center rounded text-zinc-500 transition-colors hover:bg-zinc-800 hover:text-zinc-300"
          title="Clear console output"
        >
          <Trash2 class="h-3.5 w-3.5" />
        </button>

        <!-- Maximize / Minimize -->
        <button
          type="button"
          onclick={() => (isMaximized = !isMaximized)}
          class="flex h-6 w-6 cursor-pointer items-center justify-center rounded text-zinc-500 transition-colors hover:bg-zinc-800 hover:text-zinc-300"
          title={isMaximized ? 'Restore height' : 'Maximize height'}
        >
          {#if isMaximized}
            <Minimize2 class="h-3.5 w-3.5" />
          {:else}
            <Maximize2 class="h-3.5 w-3.5" />
          {/if}
        </button>

        <!-- Close Button -->
        <button
          type="button"
          onclick={() => applicationStore.closeConsole()}
          class="flex h-6 w-6 cursor-pointer items-center justify-center rounded text-zinc-500 transition-colors hover:bg-zinc-800 hover:text-zinc-200"
          title="Close console (Cmd+J)"
        >
          <X class="h-3.5 w-3.5" />
        </button>
      </div>
    </div>

    <!-- Log Messages Container -->
    <div
      bind:this={terminalContainerEl}
      class="flex-1 overflow-y-auto p-2 font-mono text-[12px] leading-relaxed select-text"
    >
      {#if filteredLogs.length === 0}
        <div class="flex h-full flex-col items-center justify-center text-zinc-600 gap-1 select-none">
          <Terminal class="h-6 w-6 text-zinc-700" />
          <span>No console output available. Run or compile a test to view live logs.</span>
        </div>
      {:else}
        <div class="flex flex-col gap-1">
          {#each filteredLogs as log (log.id)}
            <div class="group flex flex-col rounded px-2 py-1 transition-colors hover:bg-zinc-900/60">
              <div class="flex items-start gap-2">
                <!-- Timestamp -->
                <span class="shrink-0 text-zinc-500 text-[11px]">{log.timestamp}</span>

                <!-- Level Badge -->
                {#if log.level === 'info'}
                  <span class="shrink-0 rounded border border-blue-500/30 bg-blue-500/10 px-1.5 py-0.2 text-[10px] font-semibold text-blue-400">
                    INFO
                  </span>
                {:else if log.level === 'step'}
                  <span class="shrink-0 rounded border border-purple-500/30 bg-purple-500/10 px-1.5 py-0.2 text-[10px] font-semibold text-purple-400">
                    STEP
                  </span>
                {:else if log.level === 'success'}
                  <span class="shrink-0 rounded border border-emerald-500/30 bg-emerald-500/10 px-1.5 py-0.2 text-[10px] font-semibold text-emerald-400">
                    SUCCESS
                  </span>
                {:else if log.level === 'warn'}
                  <span class="shrink-0 rounded border border-amber-500/30 bg-amber-500/10 px-1.5 py-0.2 text-[10px] font-semibold text-amber-400">
                    WARN
                  </span>
                {:else if log.level === 'error'}
                  <span class="shrink-0 rounded border border-red-500/30 bg-red-500/10 px-1.5 py-0.2 text-[10px] font-semibold text-red-400">
                    ERROR
                  </span>
                {/if}

                <!-- Category Badge if present -->
                {#if log.category}
                  <span class="shrink-0 text-[10px] uppercase tracking-wider text-zinc-500 font-semibold">
                    [{log.category}]
                  </span>
                {/if}

                <!-- Action Node Badge if present -->
                {#if log.actionId}
                  <button
                    type="button"
                    onclick={() => handleNodeClick(log.actionId)}
                    class="shrink-0 cursor-pointer rounded border border-blue-900/40 bg-blue-950/40 px-1.5 py-0.2 text-[10px] font-medium text-blue-300 hover:bg-blue-900/50 hover:underline"
                    title="Click to view node in panel"
                  >
                    node:{log.actionId}
                  </button>
                {/if}

                <!-- Message Body -->
                <span class="flex-1 text-zinc-300 break-words {log.level === 'error' ? 'text-red-300 font-medium' : ''}">
                  {log.message}
                </span>

                <!-- Details Expand Toggle -->
                {#if log.details}
                  <button
                    type="button"
                    onclick={() => toggleDetails(log.id)}
                    class="flex items-center gap-1 shrink-0 text-[11px] text-zinc-400 hover:text-zinc-200 transition-colors"
                  >
                    {#if expandedDetails[log.id]}
                      <ChevronDown class="h-3.5 w-3.5" />
                      <span>Hide details</span>
                    {:else}
                      <ChevronRight class="h-3.5 w-3.5" />
                      <span>Details</span>
                    {/if}
                  </button>
                {/if}
              </div>

              <!-- Expanded Details Box -->
              {#if log.details && expandedDetails[log.id]}
                <div class="mt-1.5 ml-14 overflow-x-auto rounded border border-zinc-800 bg-[#161b22] p-2 text-[11px] text-zinc-300 whitespace-pre-wrap">
                  {log.details}
                </div>
              {/if}
            </div>
          {/each}
        </div>
      {/if}
    </div>
  </div>
{/if}
