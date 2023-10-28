import { ProjectRepository, OrganizationRepository } from '$lib/server';
import type { Actions } from './$types';
import { redirect, error } from '@sveltejs/kit';
import type { Session } from '@auth/core/types';

export const actions: Actions = {
	default: async ({ request, params, locals }) => {
		const session = (await locals.getSession()) as Session;
		const data = await request.formData();
		const name = data.get('name') as string;
		const organizations = await OrganizationRepository.findByMember(session.user.id);
		const selectedOrganization = organizations.find((org) => org.slug === params.slug);
		if (!selectedOrganization) {
			throw error(404);
		}
		const organizationId = selectedOrganization.id;
		await ProjectRepository.create({ name, organizationId });
		// TODO: redirect user to the created project
		throw redirect(303, `/`);
	}
};
