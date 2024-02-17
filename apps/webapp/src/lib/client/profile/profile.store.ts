import { writable } from 'svelte/store';
import type { Organization, User } from '$lib/types';

export interface Profile extends User {
	organizations: Organization[];
}

/**
 * Svelte store for visible toast notifications
 */
export const profile = writable<Profile>();

/**
 * Set the current use session on the store
 */
export function setProfile(value: Profile | null) {
	if (value) {
		profile.set(value);
	}
}
