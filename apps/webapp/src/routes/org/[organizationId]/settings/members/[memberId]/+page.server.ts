import { getSelectedOrganization } from '$lib/server/action-utils';
import { OrganizationService } from '$lib/server/organization';
import { error, redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
	/**
	 * Remove a member from the organization.
	 */
	revokeMembership: async (event) => {
		const session = await event.locals.auth();
		if (!session || !session.user?.id) {
			return error(403);
		}
		const organization = await getSelectedOrganization(event);
		const userId = event.params.memberId;
		await OrganizationService.removeMember(organization.id, userId);
		if (session.user.id === userId) {
			redirect(303, '/');
		}
	}
};
