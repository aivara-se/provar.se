import { api } from '$lib/client/api';
import type { LayoutLoad } from './$types';

/**
 * Loads the organization's members and invitations.
 */
export const load: LayoutLoad = async (event) => {
	const { organization } = await event.parent();
	const members = await api().Organization.members(organization.id);
	const invitations = await api().Invitation.list(organization.id);
	return { members, invitations };
};
