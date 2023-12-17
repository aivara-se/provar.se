import { OrganizationRepository } from '$lib/server/organization';
import { UserRepository } from '$lib/server/user';
import { secureRandom } from '$lib/server/random';
import * as InvitationRepository from './invitation.repository';

/**
 * The length of the invitation key.
 */
const INVITATION_KEY_LENGTH = 16;

/**
 * Invites a user to join an organization.
 */
export async function invite(organizationId: string, name: string, email: string): Promise<void> {
	const existingUser = await UserRepository.findByEmail(email);
	if (existingUser) {
		const organizations = await OrganizationRepository.findByMember(existingUser.id);
		const isAlreadyMember = organizations.some((org) => org.id === organizationId);
		if (isAlreadyMember) {
			throw new Error('User is already a member of this organization');
		}
	}
	const existingInvitations = await InvitationRepository.findByOrganization(organizationId);
	const isAlreadyInvited = existingInvitations.some((inv) => inv.email === email);
	if (isAlreadyInvited) {
		throw new Error('User is already invited to this organization');
	}
	const key = createInvitationKey();
	await InvitationRepository.create({ key, name, email, organizationId });
	// TODO: Send an email to the user with a link to the invitation key.
}

/**
 * Checks whether the user can accept the invitation.
 */
export async function canAccept(key: string, userId: string): Promise<boolean> {
	const user = await UserRepository.findById(userId);
	const invitation = await InvitationRepository.findByKey(key);
	if (!invitation || !user) {
		return false;
	}
	if (invitation.email !== user.email) {
		return false;
	}
	const organization = await OrganizationRepository.findById(invitation.organizationId);
	if (!organization) {
		return false;
	}
	if (organization.members.includes(userId)) {
		return false;
	}
	return true;
}

/**
 * Accepts an invitation to join an organization.
 */
export async function accept(key: string, userId: string): Promise<void> {
	const invitation = await InvitationRepository.findByKey(key);
	if (!invitation) {
		throw new Error('Invitation not found');
	}
	if (!canAccept(key, userId)) {
		throw new Error('User cannot accept this invitation');
	}
	await Promise.all([
		OrganizationRepository.addMember(invitation.organizationId, userId),
		InvitationRepository.remove(invitation.organizationId, invitation.id)
	]);
}

/**
 * Generates a secure invitation key.
 */
export function createInvitationKey() {
	return secureRandom(INVITATION_KEY_LENGTH);
}
