import { InvitationRepository, InvitationService } from '$lib/server/invitation';
import { OrganizationRepository } from '$lib/server/organization';
import type { Session } from '@auth/core/types';
import { error, redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const session = await event.locals.getSession();
	if (!session || !session.user?.id) {
		redirect(302, `/auth/login`);
	}
	const invitationKey = event.params.invitationKey;
	const invitation = await InvitationRepository.findByKey(invitationKey);
	if (!invitation) {
		error(404, 'Invitation not found');
	}
	const canAccept = await InvitationService.canAccept(invitationKey, session.user?.id);
	if (!canAccept) {
		error(403, 'User cannot accept this invitation');
	}
	const organization = await OrganizationRepository.findById(invitation.organizationId);
	return { invitation, organizationName: organization?.name };
};

export const actions: Actions = {
	default: async (event) => {
		if (!event.params.invitationKey) {
			error(404);
		}
		const session = (await event.locals.getSession()) as Session;
		const invitation = await InvitationRepository.findByKey(event.params.invitationKey);
		if (!invitation) {
			throw new Error('Invitation not found');
		}
		await InvitationService.accept(event.params.invitationKey, session.user.id);
		redirect(303, `/org/${invitation.organizationId}`);
	}
};
