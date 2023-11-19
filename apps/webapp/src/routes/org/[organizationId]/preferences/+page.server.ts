import { CredentialRepository, OrganizationRepository, getSelectedOrganization } from '$lib/server';
import type { Actions } from './$types';

export const actions: Actions = {
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
			key: secureRandom(64)
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
		await CredentialRepository.revoke(organization.id, id);
	}
};

/**
 * Generates a secure random string.
 */
function secureRandom(count: number): string {
	const array = new Uint8Array(count);
	crypto.getRandomValues(array);
	return Buffer.from(array).toString('base64url');
}
