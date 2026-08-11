<script lang="ts">
  import { applicationStore } from '../../stores/application-store.svelte';
  import { settingsStore } from '../../stores/settings-store.svelte';
  import {
    SettingsService,
    SETTINGS_SCHEMA,
    validateFlatSettings,
    unflattenSettings,
  } from '../../services/settings-service';
  import { X, Eye, EyeOff } from 'lucide-svelte';

  let flatValues = $state<Record<string, any>>({});
  let showPassword = $state<Record<string, boolean>>({});
  let errors = $state<Record<string, string>>({});
  let saving = $state(false);

  $effect(() => {
    if (applicationStore.modalKind === 'settings') {
      void (async () => {
        try {
          flatValues = await SettingsService.getFlatSettings();
          errors = {};
          showPassword = {};
        } catch (e) {
          console.error('SettingsModal: load settings failed:', e);
        }
      })();
    }
  });

  $effect(() => {
    if (applicationStore.modalKind !== 'settings') return;
    const handleKeydown = (e: KeyboardEvent) => {
      if (e.key === 'Escape') {
        applicationStore.closeModal();
      }
    };
    window.addEventListener('keydown', handleKeydown);
    return () => window.removeEventListener('keydown', handleKeydown);
  });

  async function save() {
    errors = validateFlatSettings(flatValues);
    if (Object.keys(errors).length > 0) {
      return;
    }

    saving = true;
    try {
      await SettingsService.saveFlatSettings(flatValues);
      const updatedDomainSettings = unflattenSettings(flatValues);
      settingsStore.settings = updatedDomainSettings;
      settingsStore.dismissSetupWizard();
      applicationStore.closeModal();
    } catch (e) {
      console.error('SettingsModal: save settings failed:', e);
    } finally {
      saving = false;
    }
  }

  function togglePassword(id: string) {
    showPassword[id] = !showPassword[id];
  }
</script>

{#if applicationStore.modalKind === 'settings'}
  <div
    class="fixed inset-0 z-[200] flex flex-col bg-[#0d1117] text-zinc-100"
    role="dialog"
  >
    <!-- Top Header Bar -->
    <header class="flex h-14 shrink-0 items-center justify-end border-b border-zinc-800 bg-[#161b22] px-6">
      <div class="flex items-center gap-3">
        <button
          type="button"
          class="rounded-md px-3.5 py-1.5 text-xs font-medium text-zinc-400 transition-colors hover:bg-zinc-800 hover:text-zinc-200"
          onclick={() => applicationStore.closeModal()}
        >
          Cancel
        </button>
        <button
          type="button"
          disabled={saving}
          class="rounded-md bg-indigo-600 px-4 py-1.5 text-xs font-medium text-white transition-colors hover:bg-indigo-500 disabled:opacity-50"
          onclick={save}
        >
          {saving ? 'Saving...' : 'Save Settings'}
        </button>
        <button
          type="button"
          class="ml-2 rounded-md p-1.5 text-zinc-400 transition-colors hover:bg-zinc-800 hover:text-zinc-200"
          title="Close Settings"
          onclick={() => applicationStore.closeModal()}
        >
          <X class="h-4 w-4" />
        </button>
      </div>
    </header>

    <!-- Main Content Workspace (Flat Settings List) -->
    <div class="flex-1 overflow-y-auto p-8">
      <div class="mx-auto max-w-2xl space-y-5">
        {#each SETTINGS_SCHEMA as def (def.id)}
          <div>
            <label class="mb-1 block text-xs text-zinc-500" for={def.id}>
              {def.label}
            </label>

            {#if def.type === 'select'}
              <select
                id={def.id}
                bind:value={flatValues[def.id]}
                class="h-8 w-full rounded border border-zinc-700/50 bg-[#21262d] px-2.5 text-xs text-zinc-200 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 focus:outline-none"
              >
                {#each def.options ?? [] as opt}
                  <option value={opt.value}>{opt.label}</option>
                {/each}
              </select>
            {:else if def.type === 'password'}
              <div class="relative">
                <input
                  id={def.id}
                  type={showPassword[def.id] ? 'text' : 'password'}
                  bind:value={flatValues[def.id]}
                  placeholder={def.placeholder}
                  class="h-8 w-full rounded border border-zinc-700/50 bg-[#21262d] px-2.5 pr-10 font-mono text-xs text-zinc-200 placeholder-zinc-500 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 focus:outline-none"
                />
                <button
                  type="button"
                  class="absolute right-2 top-1/2 -translate-y-1/2 rounded p-1 text-zinc-400 hover:text-zinc-200"
                  title={showPassword[def.id] ? 'Hide password' : 'Show password'}
                  onclick={() => togglePassword(def.id)}
                >
                  {#if showPassword[def.id]}
                    <EyeOff class="h-3.5 w-3.5" />
                  {:else}
                    <Eye class="h-3.5 w-3.5" />
                  {/if}
                </button>
              </div>
            {:else}
              <input
                id={def.id}
                type="text"
                bind:value={flatValues[def.id]}
                placeholder={def.placeholder}
                class="h-8 w-full rounded border border-zinc-700/50 bg-[#21262d] px-2.5 text-xs text-zinc-200 placeholder-zinc-500 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 focus:outline-none"
              />
            {/if}

            {#if def.helpText}
              <p class="mt-1 text-[11px] text-zinc-500">{def.helpText}</p>
            {/if}

            {#if errors[def.id]}
              <p class="mt-1 text-[11px] text-red-400">{errors[def.id]}</p>
            {/if}
          </div>
        {/each}
      </div>
    </div>
  </div>
{/if}