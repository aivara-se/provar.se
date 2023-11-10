import { page } from '$app/stores';
import { derived } from 'svelte/store';
import type { Organization } from '../types';

/**
 * This is a derived store that returns the selected organization from the
 * current page's data. Returns null if outside an organization page.
 */
export const selectedOrg = derived(page, ($page) => {
	const organizationId = $page.params.organizationId;
	const orgs: Organization[] = $page.data.organizations;
	if (!organizationId || !orgs) {
		return null;
	}
	return orgs.find((org) => org.id === organizationId);
});
