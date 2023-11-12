import { ProjectRepository, getSelectedOrganization } from '$lib/server';
import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
	default: async (event) => {
		const organization = await getSelectedOrganization(event);
		const data = await event.request.formData();
		const name = data.get('name') as string;
		const projectId = await ProjectRepository.create({ name, organizationId: organization.id });
		throw redirect(303, `/org/${organization.id}/feedback/project/${projectId}`);
	}
};
