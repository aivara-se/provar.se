import * as OrganizationRepository from '../organization/organization.repository';
import * as UserRepository from '../user/user.repository';
import * as InvitationRepository from './invitation.repository';

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
