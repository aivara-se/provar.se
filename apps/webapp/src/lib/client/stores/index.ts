import type { Organization, User } from '$lib/client/types';
import { writable } from 'svelte/store';

/**
 * Svelte store for users organizations list
 */
export const organizations = writable<Organization[]>([]);

/**
 * Svelte store for current logged in user
 */
export const user = writable<User | null>(null);
