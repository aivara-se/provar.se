<script lang="ts">
  import { AlertCircle, AlertTriangle, ExternalLink, ShieldCheck } from 'lucide-svelte';
  import { editorStore } from '../../stores/editor-store.svelte';
  import { applicationStore } from '../../stores/application-store.svelte';
  import PanelHeader from './PanelHeader.svelte';

  function focusNode(nodeId?: string) {
    if (!nodeId) return;
    editorStore.selectedNodeId = nodeId;
    applicationStore.openNodePanel();
  }

  let totalIssues = $derived(
    editorStore.diagnostics.errors.length + editorStore.diagnostics.warnings.length
  );
</script>

<div class="flex h-full flex-col">
  <PanelHeader title="Graph Diagnostics">
    {#if totalIssues > 0}
      <span class="rounded-full bg-amber-500/20 px-2 py-0.5 text-xs font-semibold text-amber-300">
        {totalIssues} {totalIssues === 1 ? 'issue' : 'issues'}
      </span>
    {/if}
  </PanelHeader>

  <div class="flex-1 space-y-6 overflow-y-auto p-6 text-xs">
    {#if totalIssues === 0}
      <div class="flex flex-col items-center justify-center py-12 text-center text-zinc-400">
        <ShieldCheck class="mb-3 h-10 w-10 text-emerald-400" />
        <p class="text-sm font-medium text-zinc-200">No Graph Validation Issues</p>
        <p class="mt-1 text-xs text-zinc-500">Your test graph is clean, valid, and ready to run.</p>
      </div>
    {:else}
      <div class="space-y-2.5">
        {#each editorStore.diagnostics.errors as diag (diag.id)}
          <div class="group rounded-lg border border-red-900/40 bg-red-950/20 p-3 transition-colors hover:border-red-700/60">
            <div class="flex items-center justify-between gap-2">
              <span class="rounded bg-red-500/20 px-1.5 py-0.5 text-[10px] font-mono font-bold tracking-wide text-red-300">
                {diag.code}
              </span>

              {#if diag.nodeId}
                <button
                  type="button"
                  onclick={() => focusNode(diag.nodeId)}
                  class="flex items-center gap-1 text-[11px] font-medium text-red-300 hover:text-red-100 hover:underline cursor-pointer"
                  title="Focus action node on canvas"
                >
                  <span>Node: {diag.nodeId}</span>
                  <ExternalLink class="h-3 w-3" />
                </button>
              {/if}
            </div>

            <p class="mt-2 text-xs leading-relaxed text-zinc-300">{diag.message}</p>
          </div>
        {/each}

        {#each editorStore.diagnostics.warnings as diag (diag.id)}
          <div class="group rounded-lg border border-amber-900/40 bg-amber-950/20 p-3 transition-colors hover:border-amber-700/60">
            <div class="flex items-center justify-between gap-2">
              <span class="rounded bg-amber-500/20 px-1.5 py-0.5 text-[10px] font-mono font-bold tracking-wide text-amber-300">
                {diag.code}
              </span>

              {#if diag.nodeId}
                <button
                  type="button"
                  onclick={() => focusNode(diag.nodeId)}
                  class="flex items-center gap-1 text-[11px] font-medium text-amber-300 hover:text-amber-100 hover:underline cursor-pointer"
                  title="Focus action node on canvas"
                >
                  <span>Node: {diag.nodeId}</span>
                  <ExternalLink class="h-3 w-3" />
                </button>
              {/if}
            </div>

            <p class="mt-2 text-xs leading-relaxed text-zinc-300">{diag.message}</p>
          </div>
        {/each}
      </div>
    {/if}
  </div>
</div>
