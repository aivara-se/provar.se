import { getSelectedOrganization } from '$lib/server/action-utils';
import { OrganizationService } from '$lib/server/organization';
import type { Session } from '@auth/core/types';
import type { Actions } from './$types';
import { redirect } from '@sveltejs/kit';

export const actions: Actions = {
	/**
	 * Remove a member from the organization.
	 */
	revokeMembership: async (event) => {
		const session = (await event.locals.auth()) as Session;
		const organization = await getSelectedOrganization(event);
		const userId = event.params.memberId;
		await OrganizationService.removeMember(organization.id, userId);
		if (session.user.id === userId) {
			redirect(303, '/');
		}
	}
};
