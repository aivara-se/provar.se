import { redirect } from '@sveltejs/kit';
import { getOrganizations } from '$lib/api';
import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	const organizations = await getOrganizations();
	if (organizations.length === 0) {
		throw redirect(302, `/org/create`);
	}
	if (organizations.length === 1) {
		throw redirect(302, `/org/${organizations[0].id}`);
	}
	return { organizations };
};
