import { getSelectedOrganization } from '$lib/server/action-utils';
import { InvitationService } from '$lib/server/invitation';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import type { Actions, PageServerLoad } from './$types';
import { schema } from './(forms)/InviteMemberForm.schema';

export const load: PageServerLoad = async () => {
	const form = await superValidate(zod(schema));
	return { InviteMemberForm: form };
};

export const actions: Actions = {
	/**
	 * Invite a new member to join the organization.
	 */
	inviteMember: async (event) => {
		const organization = await getSelectedOrganization(event);
		const form = await superValidate(event.request, zod(schema));
		const { name, email } = form.data;
		await InvitationService.invite(organization.id, name, email);
	}
};
