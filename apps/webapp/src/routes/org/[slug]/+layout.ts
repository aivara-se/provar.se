import { error } from '@sveltejs/kit';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async (event) => {
	const { organizations } = await event.parent();
	const organization = organizations.find((org) => org.slug === event.params.slug);
	if (!organization) {
		error(404);
	}
	return { organization };
};
