import { CredentialRepository } from '$lib/server/credential';
import { ProjectRepository } from '$lib/server/project';
import { error } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async (event) => {
	const { organizations } = await event.parent();
	const organization = organizations.find((org) => org.id === event.params.organizationId);
	if (!organization) {
		error(404);
	}
	const [credentials, projects] = await Promise.all([
		CredentialRepository.findByOrganization(organization.id),
		ProjectRepository.findByOrganization(organization.id)
	]);
	return { organization, credentials, projects };
};
