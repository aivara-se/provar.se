import { FeedbackRepository } from '$lib/server/feedback';
import { error } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async (event) => {
	const { organization } = await event.parent();
	const feedback = await FeedbackRepository.findById(organization.id, event.params.feedbackId);
	if (!feedback) {
		error(404);
	}
	return { feedback };
};
