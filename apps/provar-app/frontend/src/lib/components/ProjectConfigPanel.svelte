<script lang="ts">
  import { projectStore } from '../stores/project-store.svelte';
  import { uiStore } from '../stores/ui-store.svelte';
  import PanelHeader from './PanelHeader.svelte';

  // ProjectConfigPanel edits the project's `.provar/config.yml` via the
  // Config binding. Config.LoadConfig returns the raw map; for v1 we
  // show the `variables` block as a JSON textarea (other fields like
  // browser: width/height are advanced and stay on disk untouched).
  let variablesJson = $state('{}');
  let jsonError = $state<string | null>(null);
  let dirty = $state(false);

  // Whenever the disk-side config changes (project opened / external save),
  // we reset the local editor — but only when the user hasn't queued an
  // unsaved change of their own. `dirty` is the gate.
  $effect(() => {
    if (dirty) return;
    const cfg = projectStore.config as { variables?: unknown } | null;
    variablesJson = JSON.stringify(cfg?.variables ?? {}, null, 2);
  });

  function parseVariables(): Record<string, unknown> | null {
    try {
      const parsed = JSON.parse(variablesJson);
      if (parsed === null || typeof parsed !== 'object' || Array.isArray(parsed)) {
        jsonError = 'Variables must be a JSON object.';
        return null;
      }
      return parsed as Record<string, unknown>;
    } catch (e) {
      jsonError = (e as Error).message;
      return null;
    }
  }

  // keystroke commit
  function onInput() {
    jsonError = null;
    dirty = true;
  }

  async function save() {
    const vars = parseVariables();
    if (vars === null) return;
    jsonError = null;
    await projectStore.saveConfig({
      ...(projectStore.config ?? {}),
      variables: vars,
    });
    dirty = false;
  }
</script>

<PanelHeader title="Project Settings">
  <button
    type="button"
    class="rounded p-1 text-xs font-medium text-zinc-300 transition-colors hover:bg-zinc-800 disabled:opacity-50"
    disabled={!dirty}
    onclick={save}
  >
    Save
  </button>
</PanelHeader>

<div class="flex-1 overflow-y-auto p-6 text-xs">
  <label class="mb-1 block text-zinc-500" for="config-vars">Variables (JSON)</label>
  <textarea
    id="config-vars"
    bind:value={variablesJson}
    oninput={onInput}
    rows="10"
    spellcheck="false"
    class="w-full rounded-lg border border-zinc-700/50 bg-[#0d1117] p-3 font-mono leading-relaxed text-zinc-200 placeholder-zinc-600 focus:border-indigo-500 focus:ring-1 focus:ring-indigo-500 focus:outline-none {jsonError
      ? 'border-red-500/50 ring-1 ring-red-500/20'
      : ''}"
  ></textarea>
  {#if jsonError}
    <p class="mt-2 font-mono text-[10px] leading-tight text-red-400/80">
      {jsonError}
    </p>
  {/if}
  <p class="mt-3 text-zinc-500">
    Variables are substituted into compiled Lua at run time.
    Use <span class="font-mono text-zinc-400">{'{{name}}'}</span> in titles and
    info fields; the engine will replace it with this map's value.
  </p>
</div>
