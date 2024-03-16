import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = async (event) => {
	const { organizations } = await event.parent();
	if (organizations.length === 0) {
		redirect(302, `/org/create`);
	}
};
