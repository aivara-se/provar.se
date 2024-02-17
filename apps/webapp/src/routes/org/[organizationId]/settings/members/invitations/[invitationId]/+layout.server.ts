import { error } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async (event) => {
	const { invitations } = await event.parent();
	const invitation = invitations.find((inv) => inv.id === event.params.invitationId);
	if (!invitation) {
		error(404);
	}
	return { invitation };
};
