import { getSelectedOrganization } from '$lib/server/action-utils';
import { OrganizationRepository, OrganizationService } from '$lib/server/organization';
import type { Session } from '@auth/core/types';
import { redirect } from '@sveltejs/kit';
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
	},

	/**
	 * Leaves the organization and delete the organization if there are no other members.
	 */
	leaveOrganization: async (event) => {
		const session = (await event.locals.getSession()) as Session;
		const organization = await getSelectedOrganization(event);
		await OrganizationService.removeMember(organization.id, session.user.id);
		redirect(302, '/');
	},

	/**
	 * Removes the organization and all associated data.
	 */
	deleteOrganization: async (event) => {
		const organization = await getSelectedOrganization(event);
		await OrganizationService.remove(organization.id);
		redirect(302, '/');
	}
};
