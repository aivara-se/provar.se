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
	defaultLanguage: string;
	brandingColors: {
		primary: string;
		secondary: string;
	};
}

/**
 * The default config value.
 */
const defaultValue: Config = {
	configured: false,
	feedbackTypes: [],
	defaultLanguage: 'en',
	brandingColors: {
		primary: '#000000',
		secondary: '#000000'
	}
};

/**
 * The current config value.
 */
let configValue: Config = loadConfig();

/**
 * The store for the config.
 */
export const config = writable<Config>(configValue);

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
	return {
		...defaultValue,
		...JSON.parse(stored)
	};
}

/**
 * Returns the current config value in the store.
 */
export function getConfig() {
	return configValue;
}

/**
 * Resets the config store to the default value.
 */
export function resetConfig() {
	if (!browser) {
		return;
	}
	configValue = defaultValue;
	config.set(configValue);
	localStorage.removeItem(STORAGE_KEY);
}

/**
 * Updates the config store with the provided config.
 */
export function setConfig(newConfig: Partial<Config>) {
	if (!browser) {
		return;
	}
	configValue = { ...configValue, ...newConfig, configured: true };
	config.set(configValue);
	localStorage.setItem(STORAGE_KEY, JSON.stringify(configValue));
}
