import { getOrganizations } from '$lib/api';
import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	const organizations = await getOrganizations();
	return { organizations };
};
