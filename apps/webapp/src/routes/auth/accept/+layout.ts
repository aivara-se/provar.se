import { api } from '$lib/client/api';
import { redirect } from '@sveltejs/kit';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async (event) => {
	const { organizations } = await event.parent();
	const organizationId = event.url.searchParams.get('organizationId');
	if (!organizationId) {
		return { invitedOrganization: null };
	}
	const organization = organizations.find((org) => org.id === organizationId);
	if (organization) {
		redirect(302, `/org/${organization.slug}`);
	}
	const invitedOrganization = await api().Organization.public(organizationId);
	return { invitedOrganization };
};
