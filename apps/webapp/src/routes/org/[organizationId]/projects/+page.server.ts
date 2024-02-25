import { getSelectedOrganization } from '$lib/server/action-utils';
import { ProjectRepository } from '$lib/server/project';
import { redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import type { Actions, PageServerLoad } from './$types';
import { schema } from './(forms)/CreateProjectForm.schema';

export const load: PageServerLoad = async () => {
	const form = await superValidate(zod(schema));
	return { CreateProjectForm: form };
};

export const actions: Actions = {
	/**
	 * Creates a new client credential for the organization.
	 */
	createProject: async (event) => {
		const organization = await getSelectedOrganization(event);
		const form = await superValidate(event.request, zod(schema));
		const project = {
			name: form.data.name,
			organizationId: organization.id,
			feedbackType: form.data.feedbackType
		};
		const projectId = await ProjectRepository.create(project);
		redirect(302, `/org/${organization.id}/projects/${projectId}`);
	}
};
