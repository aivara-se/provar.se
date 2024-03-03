import { FeedbackRepository } from '$lib/server/feedback';
import { parseDateRange } from '$lib/shared/dates';
import { parseSearch } from '$lib/shared/search';
import type { PageServerLoad } from './$types';

const DEFAULT_LIMIT = 20;

const EMPTY_RESPONSE = {
	count: 0,
	items: [],
	pages: 0
};

/**
 * Loads feedbacks to display on the page.
 */
export const load: PageServerLoad = async (event) => {
	const { organization } = await event.parent();
	const range = {
		from: event.url.searchParams.get('beg') ?? '',
		to: event.url.searchParams.get('end') ?? ''
	};
	if (!range.from && !range.to) {
		return EMPTY_RESPONSE;
	}
	const date = parseDateRange(range);
	const page = Number.parseInt(event.url.searchParams.get('page') ?? '1');
	const search = parseSearch(event.url.searchParams.get('search') ?? '');
	const options: FeedbackRepository.FindOptions = { date, page, limit: DEFAULT_LIMIT, search };
	const [items, count] = await Promise.all([
		FeedbackRepository.findByOrganization(organization.id, options),
		FeedbackRepository.countByOrganization(organization.id, options)
	]);
	return {
		feedbacks: {
			count: count,
			pages: Math.ceil(count / DEFAULT_LIMIT),
			items: items
		}
	};
};
