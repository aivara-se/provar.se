import { CredentialRepository, CredentialService } from '$lib/server/credential';
import { FeedbackRepository } from '$lib/server/feedback';
import { ProjectRepository } from '$lib/server/project';
import * as OrganizationRepository from './organization.repository';

/**
 * Create a new organization in the database with the given information.
 * Also creates a default project and api key for the new organization.
 */
export async function create(userId: string, name: string): Promise<string> {
	const id = await OrganizationRepository.create(userId, { name });
	await CredentialRepository.create({
		organizationId: id,
		name: 'Default',
		key: CredentialService.createCredentialKey()
	});
	return id;
}

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
