import { OrganizationRepository } from '$lib/server/organization';
import type { Session } from '@auth/core/types';
import type { LayoutServerLoad } from './$types';

/**
 * This function is also called when the user is not logged in.
 */
export const load: LayoutServerLoad = async (event) => {
	const session = (await event.locals.getSession()) as Session;
	if (!session) {
		return { session: null, organizations: [] };
	}
	const organizations = await OrganizationRepository.findByMember(session.user.id);
	return { session, organizations };
};
