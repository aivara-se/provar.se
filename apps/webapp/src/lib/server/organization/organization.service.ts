import { CredentialRepository } from '$lib/server/credential';
import { FeedbackRepository } from '$lib/server/feedback';
import { ProjectRepository } from '$lib/server/project';
import * as OrganizationRepository from './organization.repository';

/**
 * Update organization in the database with the given information.
 */
export async function remove(organizationId: string): Promise<void> {
	await CredentialRepository.removeAll(organizationId);
	await FeedbackRepository.removeAll(organizationId);
	await ProjectRepository.removeAll(organizationId);
	await OrganizationRepository.remove(organizationId);
}

/**
 * Remove a member from an organization. If the organization has no members left, it will be removed.
 */
export async function removeMember(organizationId: string, userId: string): Promise<void> {
	await OrganizationRepository.removeMember(organizationId, userId);
	const organization = await OrganizationRepository.findById(organizationId);
	if (!organization) {
		return;
	}
	if (organization.members.length === 0) {
		await remove(organizationId);
	}
}
