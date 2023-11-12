import { OrganizationRepository } from '$lib/server';
import type { Session } from '@auth/core/types';
import { error, type RequestEvent } from '@sveltejs/kit';

/**
 * A helper function to get the selected organization using path params.
 * Expects "organizationId" to be present in the path params.
 */
export async function getSelectedOrganization(event: RequestEvent) {
	if (!event.params.organizationId) {
		throw error(404);
	}
	const session = (await event.locals.getSession()) as Session;
	const organization = await OrganizationRepository.findById(event.params.organizationId);
	if (!organization || !organization?.members.includes(session.user.id)) {
		throw error(404);
	}
	return organization;
}
