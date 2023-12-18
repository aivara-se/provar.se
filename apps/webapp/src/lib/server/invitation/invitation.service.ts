import { PUBLIC_PROVAR_APP_URL } from '$env/static/public';
import { OrganizationRepository } from '$lib/server/organization';
import { secureRandom } from '$lib/server/random';
import { UserRepository } from '$lib/server/user';
import { EmailService, InvitationTemplate } from '../email-utils';
import * as InvitationRepository from './invitation.repository';

/**
 * The length of the invitation key.
 */
const INVITATION_KEY_LENGTH = 16;

/**
 * Invites a user to join an organization.
 */
export async function invite(organizationId: string, name: string, email: string): Promise<void> {
	const organization = await OrganizationRepository.findById(organizationId);
	if (!organization) {
		throw new Error('Organization not found');
	}
	const existingUser = await UserRepository.findByEmail(email);
	if (existingUser) {
		const isAlreadyMember = organization.members.some((id) => id === existingUser.id);
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
	await sendInvitationEmail(name, email, key, organization.name);
}

/**
 * Resends an invitation email.
 */
export async function resend(organizationId: string, id: string): Promise<void> {
	const invitation = await InvitationRepository.findById(id);
	if (!invitation) {
		throw new Error('Invitation not found');
	}
	const organization = await OrganizationRepository.findById(organizationId);
	if (!organization) {
		throw new Error('Organization not found');
	}
	if (invitation.organizationId !== organizationId) {
		throw new Error('Invitation does not belong to this organization');
	}
	const { name, email, key } = invitation;
	await sendInvitationEmail(name, email, key, organization.name);
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
	if (!(await canAccept(key, userId))) {
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

/**
 * Creates a link that can be used to accept an invitation.
 */
function createInvitationLink(key: string): string {
	return `${PUBLIC_PROVAR_APP_URL}/auth/accept/${key}`;
}

/**
 * Sends an invitation email with a link to accept the invite.
 */
async function sendInvitationEmail(
	name: string,
	email: string,
	key: string,
	organization: string
): Promise<void> {
	await EmailService.send({
		toEmail: email,
		template: InvitationTemplate,
		options: { name: name, link: createInvitationLink(key), organization }
	});
}
