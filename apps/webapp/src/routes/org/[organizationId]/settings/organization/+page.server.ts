import { getSelectedOrganization } from '$lib/server/action-utils';
import { OrganizationRepository } from '$lib/server/organization';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import type { Actions, PageServerLoad } from './$types';
import { schema } from './(forms)/UpdateOrganizationForm.schema';

export const load: PageServerLoad = async (event) => {
	const { organization } = await event.parent();
	const formData = { name: organization.name, description: organization.description };
	const form = await superValidate(formData, zod(schema));
	return { UpdateOrganizationForm: form };
};

export const actions: Actions = {
	/**
	 * Updates the organization's name and other details.
	 */
	updateOrganization: async (event) => {
		const organization = await getSelectedOrganization(event);
		const form = await superValidate(event.request, zod(schema));
		const { name, description } = form.data;
		await OrganizationRepository.update(organization.id, { name, description });
	}
};
