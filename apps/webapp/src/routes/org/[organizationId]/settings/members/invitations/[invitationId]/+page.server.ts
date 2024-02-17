import { getSelectedOrganization } from '$lib/server/action-utils';
import { InvitationRepository, InvitationService } from '$lib/server/invitation';
import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
	/**
	 * Resends an invitation email.
	 */
	resendInvitation: async (event) => {
		const organization = await getSelectedOrganization(event);
		const invitationId = event.params.invitationId;
		await InvitationService.resend(organization.id, invitationId);
	},

	/**
	 * Revokes an invitation to join the organization.
	 */
	revokeInvitation: async (event) => {
		const organization = await getSelectedOrganization(event);
		const invitationId = event.params.invitationId;
		await InvitationRepository.remove(organization.id, invitationId);
		redirect(302, `/org/${organization.id}/settings/members`);
	}
};
