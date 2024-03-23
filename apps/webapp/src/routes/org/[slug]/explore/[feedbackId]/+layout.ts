import { error } from '@sveltejs/kit';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async (event) => {
	const { organization } = await event.parent();
	// TODO: Load feedback from the database.
	const feedback = { id: event.params.feedbackId, organizationId: organization.id };
	if (!feedback) {
		error(404);
	}
	return { feedback };
};
