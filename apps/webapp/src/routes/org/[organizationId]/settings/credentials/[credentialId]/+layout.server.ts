import { error } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

/**
 * Loads data required for all tabs on organization preferences page.
 */
export const load: LayoutServerLoad = async (event) => {
	const { credentials } = await event.parent();
	const credential = credentials.find((cred) => cred.id === event.params.credentialId);
	if (!credential) {
		error(403);
	}
	return { credential };
};
