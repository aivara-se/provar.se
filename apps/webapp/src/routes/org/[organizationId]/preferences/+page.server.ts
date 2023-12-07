import {
	CredentialRepository,
	CredentialService,
	OrganizationRepository,
	OrganizationService,
	getSelectedOrganization
} from '$lib/server';
import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
	/**
	 * Removes the organization and all associated data.
	 */
	deleteOrganization: async (event) => {
		const organization = await getSelectedOrganization(event);
		await OrganizationService.remove(organization.id);
		throw redirect(302, '/');
	},

	/**
	 * Updates the organization's name and other details.
	 */
	updateOrganization: async (event) => {
		const organization = await getSelectedOrganization(event);
		const data = await event.request.formData();
		const name = String(data.get('name'));
		const description = String(data.get('description'));
		await OrganizationRepository.update(organization.id, { name, description });
	},

	/**
	 * Creates a new client credential for the organization.
	 */
	createCredential: async (event) => {
		const organization = await getSelectedOrganization(event);
		const data = await event.request.formData();
		const cred = {
			name: String(data.get('name')),
			organizationId: organization.id,
			key: CredentialService.createCredentialKey()
		};
		const id = await CredentialRepository.create(cred);
		return { ...cred, id };
	},

	/**
	 * Creates a new client credential for the organization.
	 */
	revokeCredential: async (event) => {
		const organization = await getSelectedOrganization(event);
		const data = await event.request.formData();
		const id = String(data.get('id'));
		await CredentialRepository.remove(organization.id, id);
	}
};
