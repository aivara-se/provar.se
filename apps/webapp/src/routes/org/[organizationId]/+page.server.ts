import { CredentialRepository, OrganizationRepository, type Organization } from '$lib/server';
import type { Session } from '@auth/core/types';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

/**
 * Loads data required for all tabs on organization preferences page.
 */
export const load: PageServerLoad = async ({ locals, params }) => {
	const session = (await locals.getSession()) as Session;
	const organization = await OrganizationRepository.findById(params.organizationId);
	if (!organization || isMember(organization, session.user.id)) {
		throw error(403);
	}
	const credentials = await CredentialRepository.findByOrganization(organization.id);
	return { credentials: credentials };
};

/**
 * Asserts that the user is a member of the organization.
 */
function isMember(organization: Organization, userId: string): boolean {
	return !organization?.members.includes(userId);
}
