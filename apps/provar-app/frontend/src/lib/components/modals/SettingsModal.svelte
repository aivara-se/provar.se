<script lang="ts">
  import Modal from './Modal.svelte';
  import { applicationStore } from '../../stores/application-store.svelte';
  import { settingsStore } from '../../stores/settings-store.svelte';
  import { SettingsService } from '../../services/settings-service';
  import { domain } from '../../../../wailsjs/go/models';

  let provider = $state('openai');
  let apiKey = $state('');
  let saving = $state(false);

  $effect(() => {
    if (applicationStore.modalKind === 'settings') {
      void (async () => {
        try {
          const current = (await SettingsService.getSettings()) ?? new domain.Settings();
          if (current.Provider) {
            provider = current.Provider;
          }
          if (current.Providers && current.Providers[provider]) {
            apiKey = current.Providers[provider].APIKey || '';
          }
        } catch (e) {
          console.error('SettingsModal: load settings failed:', e);
        }
      })();
    }
  });

  async function save() {
    saving = true;
    try {
      const current = (await SettingsService.getSettings()) ?? new domain.Settings();
      current.Provider = provider;
      if (!current.Providers) current.Providers = {};
      if (!current.Providers[provider]) {
        current.Providers[provider] = {
          Model: '',
          APIKey: apiKey.trim(),
          BaseURL: '',
        };
      } else {
        current.Providers[provider].APIKey = apiKey.trim();
      }
      await settingsStore.updateSettings(current);
      applicationStore.closeModal();
    } catch (e) {
      console.error('SettingsModal: save settings failed:', e);
    } finally {
      saving = false;
    }
  }
</script>

<Modal
  show={applicationStore.modalKind === 'settings'}
  title="Settings"
  primaryLabel={saving ? 'Saving...' : 'Save'}
  onPrimary={save}
  onClose={() => applicationStore.closeModal()}
>
  <div class="space-y-4 text-sm">
    <div>
      <label class="mb-1 block text-xs text-zinc-500" for="settings-provider">Provider</label>
      <select
        id="settings-provider"
        bind:value={provider}
        class="w-full rounded-lg border border-zinc-700/50 bg-[#21262d] p-2.5 text-zinc-200 focus:border-indigo-500 focus:ring-1 focus:ring-indigo-500 focus:outline-none"
      >
        <option value="openai">OpenAI</option>
        <option value="anthropic">Anthropic</option>
        <option value="google">Google</option>
      </select>
    </div>
    <div>
      <label class="mb-1 block text-xs text-zinc-500" for="settings-key">API Key</label>
      <input
        id="settings-key"
        type="password"
        bind:value={apiKey}
        placeholder="sk-…"
        class="w-full rounded-lg border border-zinc-700/50 bg-[#21262d] p-2.5 font-mono text-zinc-200 placeholder-zinc-500 focus:border-indigo-500 focus:ring-1 focus:ring-indigo-500 focus:outline-none"
      />
    </div>
  </div>
</Modal>