import { FeedbackRepository } from '$lib/server';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const { organization } = await event.parent();
	const feedbacks = await FeedbackRepository.findByOrganization(organization.id);
	return { feedbacks };
};
