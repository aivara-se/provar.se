import { page } from '$app/stores';
import { derived } from 'svelte/store';

export const selectedOrg = derived(page, ($page) => {
	const slug = $page.params.slug;
	const orgs = $page.data.organizations;
	if (!slug || !orgs) {
		return null;
	}
	// TODO: use strict typing when types are available for frontend
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	return orgs.find((org: any) => org.slug === slug);
});
