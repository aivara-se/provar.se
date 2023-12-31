import { getSelectedOrganization } from '$lib/server/action-utils';
import { ProjectRepository } from '$lib/server/project';
import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
	createProject: async (event) => {
		const organization = await getSelectedOrganization(event);
		const data = await event.request.formData();
		const name = data.get('name') as string;
		const projectId = await ProjectRepository.create({ name, organizationId: organization.id });
		redirect(303, `/org/${organization.id}/project/${projectId}`);
	}
};
