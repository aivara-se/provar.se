import { writable } from 'svelte/store';
import type { Toast } from './toast.types';

/**
 * Default timeout for toast notifications
 */
export const DEFAULT_TIMEOUT = 3000;

/**
 * Svelte store for visible toast notifications
 */
export const toasts = writable<Toast[]>([]);

/**
 * Generate a unique ID for a toast notification
 */
const nextToastId = (() => {
	let id = 0;
	return () => (++id).toString();
})();

/**
 * Show a toast notification with given type and content
 */
export function toast(type: Toast['type'], content: string, duration = DEFAULT_TIMEOUT) {
	const toast: Toast = { id: nextToastId(), type, content, duration };
	toasts.update((arr) => [...arr, toast]);
	setTimeout(() => {
		toasts.update((arr) => arr.filter((t) => t !== toast));
	}, duration);
}
