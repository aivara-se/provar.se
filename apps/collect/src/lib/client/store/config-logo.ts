import { writable } from 'svelte/store';
import { browser } from '$app/environment';

/**
 * The key used to store the config in local storage.
 */
const STORAGE_KEY = 'provar:config:logo';

/**
 * The current logo url value.
 */
let logoUrlValue: string | null = loadLogoUrl();

/**
 * The store for the logo url.
 */
export const logoUrl = writable<string | null>(logoUrlValue);

/**
 * Loads the logo url from local storage when available.
 */
function loadLogoUrl(): string | null {
	if (!browser) {
		return null;
	}
	const storedValue = localStorage.getItem(STORAGE_KEY);
	if (!storedValue) {
		return null;
	}
	return storedValue;
}

/**
 * Get a data url for the logo stored in local storage.
 */
export function getLogoUrl(): string | null {
	return logoUrlValue;
}

/**
 * Resets the logo to the default value.
 */
export function resetLogoUrl() {
	if (!browser) {
		return;
	}
	logoUrlValue = null;
	localStorage.removeItem(STORAGE_KEY);
	logoUrl.set(logoUrlValue);
}

/**
 * Set the logo data and type in local storage.
 */
export function setLogoUrl(dataUrl: string) {
	if (!browser) {
		return;
	}
	logoUrlValue = dataUrl;
	localStorage.setItem(STORAGE_KEY, dataUrl);
	logoUrl.set(logoUrlValue);
}
