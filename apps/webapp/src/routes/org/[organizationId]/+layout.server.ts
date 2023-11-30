import { CredentialRepository, ProjectRepository } from '$lib/server';
import { error } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

/**
 * Loads data required for all tabs on organization preferences page.
 */
export const load: LayoutServerLoad = async (event) => {
	const { organizations } = await event.parent();
	const organization = organizations.find((org) => org.id === event.params.organizationId);
	if (!organization) {
		throw error(403);
	}
	const credentials = await CredentialRepository.findByOrganization(organization.id);
	const projects = await ProjectRepository.findByOrganization(organization.id);
	return { organization, credentials, projects };
};
