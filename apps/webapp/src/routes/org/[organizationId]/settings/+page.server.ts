import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const { organization } = await event.parent();
	redirect(302, `/org/${organization.id}/settings/organization`);
};
