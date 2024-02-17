import { getSelectedOrganization } from '$lib/server/action-utils';
import { ProjectRepository } from '$lib/server/project';
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
		const cred = {
			name: form.data.name,
			organizationId: organization.id
		};
		await ProjectRepository.create(cred);
	}
};
