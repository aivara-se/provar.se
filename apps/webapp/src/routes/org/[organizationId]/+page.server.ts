import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const { organization, credentials } = await event.parent();
	if (credentials.length === 0) {
		redirect(302, `/org/${organization.id}/onboard`);
	}
	redirect(302, `/org/${organization.id}/feedback`);
};
