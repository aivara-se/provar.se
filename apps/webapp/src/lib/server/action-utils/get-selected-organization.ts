import { OrganizationRepository } from '$lib/server/organization';
import { error, type RequestEvent } from '@sveltejs/kit';

/**
 * A helper function to get the selected organization using path params.
 * Makes sure that the user is a member of the organization.
 * Expects "organizationId" to be present in the path params.
 */
export async function getSelectedOrganization(event: RequestEvent) {
	if (!event.params.organizationId) {
		error(404);
	}
	const session = await event.locals.auth();
	if (!session || !session.user?.id) {
		error(403);
	}
	const organization = await OrganizationRepository.findById(event.params.organizationId);
	if (!organization || !organization?.members.includes(session.user.id)) {
		error(404);
	}
	return organization;
}
