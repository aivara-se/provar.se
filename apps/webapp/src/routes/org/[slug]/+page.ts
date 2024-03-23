import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = async (event) => {
	const { organization } = await event.parent();
	// TODO: redirect to `/org/${organization.slug}/onboard`
	// redirect(302, `/org/${organization.slug}/explore`);
	redirect(302, `/org/${organization.slug}/onboard`);
};
