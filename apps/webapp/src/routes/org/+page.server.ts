import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const { organizations } = await event.parent();
	if (organizations.length === 0) {
		redirect(302, `/org/create`);
	}
};
