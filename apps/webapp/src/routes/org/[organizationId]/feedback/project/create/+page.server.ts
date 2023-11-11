import { OrganizationRepository, ProjectRepository } from '$lib/server';
import type { Session } from '@auth/core/types';
import { error, redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
	default: async ({ request, params, locals }) => {
		const session = (await locals.getSession()) as Session;
		const data = await request.formData();
		const name = data.get('name') as string;
		const organizations = await OrganizationRepository.findByMember(session.user.id);
		const selectedOrganization = organizations.find((org) => org.id === params.organizationId);
		if (!selectedOrganization) {
			throw error(404);
		}
		const organizationId = selectedOrganization.id;
		const projectId = await ProjectRepository.create({ name, organizationId });
		throw redirect(303, `/org/${organizationId}/feedback/project/${projectId}`);
	}
};
