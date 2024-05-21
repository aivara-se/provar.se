import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import type { FeedbackType } from '$lib/client/types';

/**
 * The key used to store the config in local storage.
 */
const STORAGE_KEY = 'provar:config';

/**
 * Config holds the configuration for the app.
 */
interface Config {
	configured: boolean;
	feedbackTypes: FeedbackType[];
}

/**
 * The default config value.
 */
const defaultValue: Config = {
	configured: false,
	feedbackTypes: []
};

/**
 * The current config value.
 */
let configValue: Config = loadConfig();

/**
 * The store for the config.
 */
const configStore = writable<Config>(configValue);

/**
 * Loads the config from local storage when available.
 */
function loadConfig(): Config {
	if (!browser) {
		return defaultValue;
	}
	const stored = localStorage.getItem(STORAGE_KEY);
	if (!stored) {
		return defaultValue;
	}
	return JSON.parse(stored);
}

/**
 * Returns the current config value in the store.
 */
export function getConfig() {
	return configValue;
}

/**
 * Updates the config store with the provided config.
 */
export function setConfig(config: Partial<Config>) {
	configValue = { ...configValue, ...config, configured: true };
	configStore.set(configValue);
	localStorage.setItem(STORAGE_KEY, JSON.stringify(configValue));
}
