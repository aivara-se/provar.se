import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = async (event) => {
	const { organization } = await event.parent();
	redirect(302, `/org/${organization.slug}/settings/organization`);
};
