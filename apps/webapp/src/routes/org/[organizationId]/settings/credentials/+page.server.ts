import { getSelectedOrganization } from '$lib/server/action-utils';
import { CredentialRepository, CredentialService } from '$lib/server/credential';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import type { Actions, PageServerLoad } from './$types';
import { schema } from './(forms)/CreateCredentialForm.schema';

export const load: PageServerLoad = async () => {
	const form = await superValidate(zod(schema));
	return { CreateCredentialForm: form };
};

export const actions: Actions = {
	/**
	 * Creates a new client credential for the organization.
	 */
	createCredential: async (event) => {
		const organization = await getSelectedOrganization(event);
		const form = await superValidate(event.request, zod(schema));
		const cred = {
			name: form.data.name,
			organizationId: organization.id,
			key: CredentialService.createCredentialKey()
		};
		await CredentialRepository.create(cred);
	}
};
