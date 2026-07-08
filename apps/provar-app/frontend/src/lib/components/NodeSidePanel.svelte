<script lang="ts">
  import { Code, Copy, Trash2, Image as ImageIcon } from 'lucide-svelte';
  import { editorStore } from '../stores/editor-store.svelte';
  import PanelHeader from './PanelHeader.svelte';

  type View = 'description' | 'code';

  let title = $state('');
  let info = $state('');
  let data = $state('');
  let view = $state<View>('description');
  let codeCopied = $state(false);

  // Sync local form state with the selected node whenever the selection
  // changes. We reset `view` so a fresh node opens on its description by
  // default; otherwise swapping between two nodes mid-code-view would be
  // surprising.
  $effect(() => {
    const node = editorStore.selectedNode;
    const id = editorStore.selectedNodeId;
    title = node?.title ?? '';
    info = node?.info ?? '';
    data = node?.data ?? '';
    view = 'description';
    codeCopied = false;
    // Force-recompute so the `data` reactivity tracks `id` and not just
    // the node shape — selecting between two nodes with the same shape
    // would otherwise miss the reset.
    void id;
  });

  function commit() {
    const id = editorStore.selectedNodeId;
    if (!id) return;
    editorStore.updateNode(id, { title, info, data });
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
  <PanelHeader title="Task Settings">
    <button
      type="button"
      class="rounded p-1.5 text-zinc-400 transition-colors hover:bg-zinc-800 hover:text-indigo-300 {view ===
      'code'
        ? 'bg-zinc-800 text-indigo-300'
        : ''}"
      title={hasCode
        ? view === 'code'
          ? 'Show description'
          : 'Show generated code'
        : 'Run `provar compile` to generate this node\'s Lua'}
      disabled={!hasCode}
      onclick={() => (view = view === 'description' ? 'code' : 'description')}
    >
      <Code class="h-3.5 w-3.5" />
    </button>
    <button
      type="button"
      class="rounded p-1.5 text-zinc-500 transition-colors hover:bg-red-500/10 hover:text-red-400"
      title="Delete node"
      onclick={deleteNode}
    >
      <Trash2 class="h-3.5 w-3.5" />
    </button>
  </PanelHeader>

  <div class="flex-1 space-y-6 overflow-y-auto p-6 text-xs">
    {#if view === 'description'}
      <div>
        <label class="mb-1 block text-zinc-500" for="node-title">Title</label>
        <input
          id="node-title"
          type="text"
          bind:value={title}
          onblur={commit}
          class="w-full rounded border border-zinc-700/50 bg-[#21262d] px-2.5 py-1.5 text-zinc-200 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 focus:outline-none"
        />
      </div>

      <div>
        <label class="mb-1 block text-zinc-500" for="node-info">Description</label>
        <textarea
          id="node-info"
          bind:value={info}
          onblur={commit}
          rows="3"
          class="w-full rounded border border-zinc-700/50 bg-[#21262d] px-2.5 py-1.5 text-zinc-200 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 focus:outline-none"
        ></textarea>
      </div>

      <div>
        <label class="mb-1 block text-zinc-500" for="node-data">Data</label>
        <textarea
          id="node-data"
          bind:value={data}
          onblur={commit}
          rows="6"
          spellcheck="false"
          placeholder="Optional arbitrary payload — preserved verbatim in the YAML."
          class="w-full rounded border border-zinc-700/50 bg-[#21262d] px-2.5 py-1.5 font-mono text-zinc-200 placeholder-zinc-600 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 focus:outline-none"
        ></textarea>
      </div>
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
          The compiled Lua body for this task. Reads from the engine's compiled
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
{/if}
