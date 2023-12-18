import { getSelectedOrganization } from '$lib/server/action-utils';
import { CredentialRepository, CredentialService } from '$lib/server/credential';
import { InvitationRepository, InvitationService } from '$lib/server/invitation';
import { OrganizationRepository, OrganizationService } from '$lib/server/organization';
import { UserRepository } from '$lib/server/user';
import type { Session } from '@auth/core/types';
import { redirect } from '@sveltejs/kit';
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
	 * Removes the organization and all associated data.
	 */
	deleteOrganization: async (event) => {
		const organization = await getSelectedOrganization(event);
		await OrganizationService.remove(organization.id);
		redirect(302, '/');
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
	},

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
