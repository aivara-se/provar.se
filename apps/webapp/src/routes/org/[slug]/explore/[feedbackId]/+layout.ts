import { api } from '$lib/client/api';
import { error } from '@sveltejs/kit';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async (event) => {
	const { organization } = await event.parent();
	const feedback = await api().Feedback.details(organization.id, event.params.feedbackId);
	if (!feedback) {
		error(404);
	}
	return { feedback };
};
