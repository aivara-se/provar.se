import { writable } from 'svelte/store';
import { createId } from '$lib/shared/utils';
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
 * Show a toast notification with given type and content
 */
export function toast(type: Toast['type'], content: string, duration = DEFAULT_TIMEOUT) {
	const toast: Toast = { id: createId(), type, content, duration };
	toasts.update((arr) => [...arr, toast]);
	setTimeout(() => {
		toasts.update((arr) => arr.filter((t) => t !== toast));
	}, duration);
}
