import { getSelectedOrganization } from '$lib/server/action-utils';
import { InvitationRepository, InvitationService } from '$lib/server/invitation';
import { OrganizationService } from '$lib/server/organization';
import { UserRepository } from '$lib/server/user';
import type { Actions, PageServerLoad } from './$types';

/**
 * Loads the organization's members.
 */
export const load: PageServerLoad = async (event) => {
	const { organization } = await event.parent();
	const members = await UserRepository.findByIds(organization.members);
	const invitations = await InvitationRepository.findByOrganization(organization.id);
	return { members, invitations };
};

export const actions: Actions = {
	/**
	 * Invite a new member to join the organization.
	 */
	inviteMember: async (event) => {
		const organization = await getSelectedOrganization(event);
		const data = await event.request.formData();
		const name = String(data.get('name'));
		const email = String(data.get('email'));
		await InvitationService.invite(organization.id, name, email);
	},

	/**
	 * Resends an invitation email.
	 */
	resendInvitation: async (event) => {
		const organization = await getSelectedOrganization(event);
		const data = await event.request.formData();
		const invitationId = String(data.get('id'));
		await InvitationService.resend(organization.id, invitationId);
	},

	/**
	 * Revokes an invitation to join the organization.
	 */
	revokeInvitation: async (event) => {
		const organization = await getSelectedOrganization(event);
		const data = await event.request.formData();
		const invitationId = String(data.get('id'));
		await InvitationRepository.remove(organization.id, invitationId);
	},

	/**
	 * Remove a member from the organization.
	 */
	revokeMembership: async (event) => {
		const organization = await getSelectedOrganization(event);
		const data = await event.request.formData();
		const userId = String(data.get('id'));
		await OrganizationService.removeMember(organization.id, userId);
	}
};
