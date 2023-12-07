import * as CredentialRepository from '../credential/credential.repository';
import * as FeedbackRepository from '../feedback/feedback.repository';
import * as ProjectRepository from '../project/project.repository';
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
