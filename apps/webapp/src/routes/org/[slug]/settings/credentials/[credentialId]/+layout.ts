import { error } from '@sveltejs/kit';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async (event) => {
	const { credentials } = await event.parent();
	const credential = credentials.find((cred) => cred.id === event.params.credentialId);
	if (!credential) {
		error(404);
	}
	return { credential };
};
