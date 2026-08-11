<script lang="ts">
  import { Code, Copy, Trash2, Image as ImageIcon, AlertTriangle, AlertCircle } from 'lucide-svelte';
  import { editorStore } from '../../stores/editor-store.svelte';
  import PanelHeader from './PanelHeader.svelte';

  type View = 'info' | 'code';

  let name = $state('');
  let info = $state('');
  let view = $state<View>('info');
  let codeCopied = $state(false);

  let nodeDiagnostics = $derived.by(() => {
    const id = editorStore.selectedNodeId;
    if (!id) return [];
    const errors = editorStore.diagnostics.errors.filter((e) => e.nodeId === id);
    const warnings = editorStore.diagnostics.warnings.filter((w) => w.nodeId === id);
    return [...errors, ...warnings];
  });

  // Sync local form state with the selected node whenever the selection
  // changes. We reset `view` so a fresh node opens on its info by
  // default.
  $effect(() => {
    const node = editorStore.selectedNode;
    const id = editorStore.selectedNodeId;
    name = node?.name ?? '';
    info = node?.info ?? '';
    view = 'info';
    codeCopied = false;
    void id;
  });

  function commit() {
    const id = editorStore.selectedNodeId;
    if (!id) return;
    editorStore.updateNode(id, { name, info });
  }

  function deleteNode() {
    const id = editorStore.selectedNodeId;
    if (!id) return;
    editorStore.deleteNode(id);
  }

  async function copyCode() {
    const node = editorStore.selectedNode;
    if (!node?.source) return;
    try {
      await navigator.clipboard.writeText(node.source);
      codeCopied = true;
      setTimeout(() => (codeCopied = false), 1500);
    } catch (e) {
      console.error('Failed to copy generated code:', e);
    }
  }

  let hasCode = $derived(!!(editorStore.selectedNode?.source ?? '').trim());
</script>

{#if editorStore.selectedNode}
  <PanelHeader title="Action Settings">
    <button
      type="button"
      class="rounded p-1.5 text-zinc-500 transition-colors hover:bg-red-500/10 hover:text-red-400"
      title="Delete action"
      onclick={deleteNode}
    >
      <Trash2 class="h-3.5 w-3.5" />
    </button>
  </PanelHeader>

  <div class="flex-1 space-y-6 overflow-y-auto p-6 text-xs">
    {#if nodeDiagnostics.length > 0}
      <div class="space-y-2 rounded-lg border border-amber-500/30 bg-amber-950/30 p-3 text-xs">
        <div class="flex items-center gap-1.5 font-semibold text-amber-400">
          <AlertTriangle class="h-4 w-4 shrink-0 text-amber-400" />
          <span>Graph Validation Issues ({nodeDiagnostics.length})</span>
        </div>
        <div class="space-y-1.5 pt-1 text-zinc-300">
          {#each nodeDiagnostics as diag (diag.id)}
            <div class="flex flex-col gap-0.5 rounded bg-zinc-900/80 p-2.5 border border-zinc-800">
              <div class="flex items-center gap-1.5 font-mono text-[10px] font-bold {diag.severity === 'error' ? 'text-red-400' : 'text-amber-400'}">
                {#if diag.severity === 'error'}
                  <AlertCircle class="h-3 w-3 shrink-0" />
                {:else}
                  <AlertTriangle class="h-3 w-3 shrink-0" />
                {/if}
                <span>{diag.code}</span>
              </div>
              <p class="text-xs leading-relaxed text-zinc-300">{diag.message}</p>
            </div>
          {/each}
        </div>
      </div>
    {/if}

    <div>
      <label class="mb-1 block text-zinc-500" for="node-name">Name</label>
      <input
        id="node-name"
        type="text"
        bind:value={name}
        onblur={commit}
        class="w-full rounded border border-zinc-700/50 bg-[#21262d] px-2.5 py-1.5 text-zinc-200 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 focus:outline-none"
      />
    </div>

    <div>
      <div class="mb-2 flex items-center justify-between">
        <label class="block text-zinc-500" for="node-info">Info</label>
        <button
          type="button"
          class="flex items-center rounded p-1.5 text-xs transition-colors hover:bg-zinc-800 hover:text-indigo-300 disabled:cursor-not-allowed disabled:opacity-40 {view ===
          'code'
            ? 'bg-zinc-800 text-indigo-300'
            : 'text-zinc-400'}"
          title={hasCode
            ? view === 'code'
              ? 'Show info'
              : 'Show generated code'
            : 'Run `provar compile` to generate this action\'s Lua'}
          disabled={!hasCode}
          onclick={() => (view = view === 'info' ? 'code' : 'info')}
          aria-pressed={view === 'code'}
        >
          <Code class="h-3.5 w-3.5" />
        </button>
      </div>

      {#if view === 'info'}
        <textarea
          id="node-info"
          bind:value={info}
          onblur={commit}
          rows="4"
          class="w-full rounded border border-zinc-700/50 bg-[#21262d] px-2.5 py-1.5 text-zinc-200 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 focus:outline-none"
        ></textarea>
      {:else}
        {#if hasCode}
          <div class="relative">
            <button
              type="button"
              onclick={copyCode}
              class="absolute top-2 right-2 z-10 flex items-center gap-1.5 rounded-md border border-zinc-800/60 bg-zinc-900/80 px-2 py-1 text-xs text-zinc-400 backdrop-blur-sm transition-colors hover:border-zinc-700 hover:text-zinc-200"
              title="Copy generated code"
            >
              <Copy class="h-3 w-3" />
              {codeCopied ? 'Copied' : 'Copy'}
            </button>
            <pre
              class="max-h-96 overflow-auto rounded-lg border border-zinc-700/50 bg-[#0d1117] p-3 pr-16 text-xs leading-relaxed text-zinc-300"><code
                >{editorStore.selectedNode.source}</code
              ></pre>
          </div>
          <p class="mt-2 text-zinc-500">
            The compiled Lua body for this action. Reads from the engine's compiled
            output on disk — recompile the test to refresh after edits.
          </p>
        {:else}
          <div
            class="flex aspect-video flex-col items-center justify-center rounded-xl border border-dashed border-zinc-800 bg-zinc-900/30 text-zinc-500"
          >
            <ImageIcon class="mb-2 h-6 w-6 opacity-50" />
            <span class="text-sm">No generated code yet</span>
            <span class="mt-1 text-[11px] text-zinc-600">
              Run <span class="font-mono">provar compile</span> from the CLI to emit it.
            </span>
          </div>
        {/if}
      {/if}
    </div>
  </div>
{/if}



