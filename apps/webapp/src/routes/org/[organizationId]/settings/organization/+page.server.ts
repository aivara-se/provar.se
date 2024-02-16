import { getSelectedOrganization } from '$lib/server/action-utils';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { OrganizationRepository } from '$lib/server/organization';
import type { Actions, PageServerLoad } from './$types';
import { schema } from './(forms)/UpdateOrganizationForm.schema';

export const load: PageServerLoad = async (event) => {
	const { organization } = await event.parent();
	const form = await superValidate(organization, zod(schema));
	return { updateOrganizationForm: form };
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
