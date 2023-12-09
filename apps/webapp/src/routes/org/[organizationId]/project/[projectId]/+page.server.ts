import { getSelectedProject, ProjectRepository } from '$lib/server';
import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
	delete: async (event) => {
		const { organization, project } = await getSelectedProject(event);
		await ProjectRepository.remove(organization.id, project.id);
		throw redirect(303, `/org/${organization.id}/project`);
	}
};
