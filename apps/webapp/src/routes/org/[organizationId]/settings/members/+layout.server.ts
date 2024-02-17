import { InvitationRepository } from '$lib/server/invitation';
import { UserRepository } from '$lib/server/user';
import type { LayoutServerLoad } from './$types';

/**
 * Loads the organization's members and invitations.
 */
export const load: LayoutServerLoad = async (event) => {
	const { organization } = await event.parent();
	const members = await UserRepository.findByIds(organization.members);
	const invitations = await InvitationRepository.findByOrganization(organization.id);
	return { members, invitations };
};
