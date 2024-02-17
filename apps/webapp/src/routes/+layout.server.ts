import { OrganizationRepository } from '$lib/server/organization';
import { UserRepository } from '$lib/server/user';
import type { LayoutServerLoad } from './$types';

/**
 * This function is also called when the user is not logged in.
 */
export const load: LayoutServerLoad = async (event) => {
	const session = await event.locals.auth();
	if (!session || !session.user?.id) {
		return { session: null, profile: null, organizations: [] };
	}
	const [profile, organizations] = await Promise.all([
		UserRepository.findById(session.user.id),
		OrganizationRepository.findByMember(session.user.id)
	]);
	return { session, profile, organizations };
};
