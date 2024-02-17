import { getSelectedOrganization } from '$lib/server/action-utils';
import { CredentialRepository } from '$lib/server/credential';
import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
	/**
	 * Creates a new client credential for the organization.
	 */
	revokeCredential: async (event) => {
		const organization = await getSelectedOrganization(event);
		const id = event.params.credentialId;
		await CredentialRepository.remove(organization.id, id);
		redirect(302, `/org/${organization.id}/settings/credentials`);
	}
};
