import { InvitationRepository, InvitationService } from '$lib/server/invitation';
import { OrganizationRepository } from '$lib/server/organization';
import { error, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const session = await event.locals.getSession();
	if (!session || !session.user?.id) {
		throw redirect(302, `/auth/login`);
	}
	const invitationKey = event.params.invitationKey;
	const invitation = await InvitationRepository.findByKey(invitationKey);
	if (!invitation) {
		throw error(404, 'Invitation not found');
	}
	const canAccept = await InvitationService.canAccept(invitationKey, session.user?.id);
	if (!canAccept) {
		throw error(403, 'User cannot accept this invitation');
	}
	const organization = await OrganizationRepository.findById(invitation.organizationId);
	return { invitation, organizationName: organization?.name };
};
