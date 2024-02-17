import { getSelectedOrganization } from '$lib/server/action-utils';
import { ProjectRepository } from '$lib/server/project';
import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
	/**
	 * Creates a new client credential for the organization.
	 */
	removeProject: async (event) => {
		// TODO: remove project reference from all feedback
		const organization = await getSelectedOrganization(event);
		const id = event.params.projectId;
		await ProjectRepository.remove(organization.id, id);
		redirect(302, `/org/${organization.id}/projects`);
	}
};
