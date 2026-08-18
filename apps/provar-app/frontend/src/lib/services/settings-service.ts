import { Project } from './bindings';
import { domain } from '../../../wailsjs/go/models';
import { applicationStore } from '../stores/application-store.svelte';

export type SettingType = 'select' | 'text' | 'password';

export interface SelectOption {
  label: string;
  value: string;
}

export interface SettingDefinition {
  id: string;
  label: string;
  helpText?: string;
  type: SettingType;
  options?: SelectOption[];
  placeholder?: string;
  defaultValue?: any;
}

export const SETTINGS_SCHEMA: SettingDefinition[] = [
  {
    id: 'provider',
    label: 'Active LLM Provider',
    helpText: 'Select the active LLM provider used for Provar test execution.',
    type: 'select',
    options: [
      { label: 'Google', value: 'google' },
      { label: 'OpenAI', value: 'openai' },
      { label: 'Anthropic', value: 'anthropic' },
    ],
    defaultValue: 'google',
  },
  {
    id: 'providers.google.model',
    label: 'Google Model',
    helpText: 'Model identifier for Google Gemini API.',
    type: 'text',
    placeholder: 'gemini-3.7-flash',
    defaultValue: 'gemini-3.7-flash',
  },
  {
    id: 'providers.google.apiKey',
    label: 'Google API Key',
    helpText: 'API Key for authenticating with Google AI Studio / Gemini.',
    type: 'password',
    placeholder: 'AIzaSy...',
  },
  {
    id: 'providers.google.baseUrl',
    label: 'Google Base URL',
    helpText: 'Optional custom endpoint base URL for Google provider.',
    type: 'text',
    placeholder: 'https://generativelanguage.googleapis.com',
  },
  {
    id: 'providers.openai.model',
    label: 'OpenAI Model',
    helpText: 'Model identifier for OpenAI API.',
    type: 'text',
    placeholder: 'gpt-5.6-terra',
    defaultValue: 'gpt-5.6-terra',
  },
  {
    id: 'providers.openai.apiKey',
    label: 'OpenAI API Key',
    helpText: 'API Key for authenticating with OpenAI.',
    type: 'password',
    placeholder: 'sk-...',
  },
  {
    id: 'providers.openai.baseUrl',
    label: 'OpenAI Base URL',
    helpText: 'Optional custom endpoint base URL for OpenAI provider.',
    type: 'text',
    placeholder: 'https://api.openai.com/v1',
  },
  {
    id: 'providers.anthropic.model',
    label: 'Anthropic Model',
    helpText: 'Model identifier for Anthropic API.',
    type: 'text',
    placeholder: 'claude-sonnet-5',
    defaultValue: 'claude-sonnet-5',
  },
  {
    id: 'providers.anthropic.apiKey',
    label: 'Anthropic API Key',
    helpText: 'API Key for authenticating with Anthropic.',
    type: 'password',
    placeholder: 'sk-ant-...',
  },
  {
    id: 'providers.anthropic.baseUrl',
    label: 'Anthropic Base URL',
    helpText: 'Optional custom endpoint base URL for Anthropic provider.',
    type: 'text',
    placeholder: 'https://api.anthropic.com',
  },
];

/**
 * flattenSettings dynamically extracts flat dotted field values (e.g. "providers.google.model")
 * from the nested domain.Settings structure.
 */
export function flattenSettings(settings: domain.Settings): Record<string, any> {
  const flat: Record<string, any> = {};
  if (!settings) return flat;

  flat['provider'] = settings.Provider || 'google';

  const providers = settings.Providers || {};
  for (const def of SETTINGS_SCHEMA) {
    if (def.id === 'provider') continue;
    const parts = def.id.split('.');
    if (parts[0] === 'providers' && parts.length === 3) {
      const providerName = parts[1];
      const fieldName = parts[2];
      const cfg = providers[providerName] || {};
      let val = '';
      if (fieldName === 'model') val = cfg.Model || '';
      else if (fieldName === 'apiKey') val = cfg.APIKey || '';
      else if (fieldName === 'baseUrl') val = cfg.BaseURL || '';
      flat[def.id] = val || def.defaultValue || '';
    }
  }

  return flat;
}

/**
 * unflattenSettings dynamically constructs a domain.Settings object from flat dotted ID pairs.
 */
export function unflattenSettings(flat: Record<string, any>): domain.Settings {
  const s = new domain.Settings();
  s.Provider = flat['provider'] || 'google';
  s.Providers = {};

  for (const def of SETTINGS_SCHEMA) {
    if (def.id === 'provider') continue;
    const parts = def.id.split('.');
    if (parts[0] === 'providers' && parts.length === 3) {
      const providerKey = parts[1];
      const fieldKey = parts[2];

      if (!s.Providers[providerKey]) {
        s.Providers[providerKey] = new domain.ProviderConfig({
          Model: '',
          APIKey: '',
          BaseURL: '',
        });
      }

      const cfg = s.Providers[providerKey];
      const val = String(flat[def.id] ?? '').trim();
      if (fieldKey === 'model') cfg.Model = val;
      if (fieldKey === 'apiKey') cfg.APIKey = val;
      if (fieldKey === 'baseUrl') cfg.BaseURL = val;
    }
  }

  return s;
}

/**
 * validateFlatSettings checks required active provider fields and returns a map of setting ID to error message.
 */
export function validateFlatSettings(flat: Record<string, any>): Record<string, string> {
  const errors: Record<string, string> = {};
  const activeProvider = flat['provider'];

  if (!activeProvider) {
    errors['provider'] = 'Active provider is required.';
  } else {
    const apiKeyId = `providers.${activeProvider}.apiKey`;
    const apiKeyVal = String(flat[apiKeyId] ?? '').trim();
    if (!apiKeyVal) {
      errors[apiKeyId] = `API Key is required for the active provider (${activeProvider}).`;
    }
    const modelId = `providers.${activeProvider}.model`;
    const modelVal = String(flat[modelId] ?? '').trim();
    if (!modelVal) {
      errors[modelId] = `Model name is required for the active provider (${activeProvider}).`;
    }
  }

  return errors;
}

export class SettingsService {
  static async getHome(): Promise<string> {
    try {
      return await Project.Home();
    } catch (e) {
      console.error('SettingsService: Home lookup failed:', e);
      throw e;
    }
  }

  static async getSettings(): Promise<domain.Settings> {
    try {
      const settings = await Project.Settings();
      return settings ?? new domain.Settings();
    } catch (e) {
      console.error('SettingsService: getSettings failed:', e);
      throw e;
    }
  }

  static async getFlatSettings(): Promise<Record<string, any>> {
    const settings = await this.getSettings();
    return flattenSettings(settings);
  }

  static async saveSettings(settings: domain.Settings): Promise<void> {
    try {
      await Project.SaveSettings(settings);
      applicationStore.showToast('info', 'Settings saved successfully');
    } catch (e) {
      const msg = e instanceof Error ? e.message : String(e);
      applicationStore.showToast('error', `Could not save settings: ${msg}`);
      throw e;
    }
  }

  static async saveFlatSettings(flatValues: Record<string, any>): Promise<void> {
    const errors = validateFlatSettings(flatValues);
    if (Object.keys(errors).length > 0) {
      const firstError = Object.values(errors)[0];
      applicationStore.showToast('error', firstError);
      throw new Error(firstError);
    }
    const settings = unflattenSettings(flatValues);
    await this.saveSettings(settings);
  }

  static async validateSettings(): Promise<void> {
    const err = await Project.ValidateSettings();
    if (err !== null) {
      throw new Error(String(err));
    }
  }
}

