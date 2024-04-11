import { api } from '$lib/client/api';
import { parseDateRange } from '$lib/client/dates';
import { parseSearch } from '$lib/client/search';
import type { PageLoad } from './$types';

const DEFAULT_LIMIT = 10;

const EMPTY_RESPONSE = {
	count: 0,
	pages: 0,
	items: []
};

/**
 * Loads feedbacks to display on the page.
 */
export const load: PageLoad = async (event) => {
	const { organization } = await event.parent();
	const range = {
		from: event.url.searchParams.get('beg') ?? '',
		to: event.url.searchParams.get('end') ?? ''
	};
	if (!range.from && !range.to) {
		return { feedbacks: EMPTY_RESPONSE };
	}
	const date = parseDateRange(range);
	const page = Number.parseInt(event.url.searchParams.get('page') ?? '1');
	const search = parseSearch(event.url.searchParams.get('search') ?? '');
	const options = {
		pageLimit: DEFAULT_LIMIT,
		pageOffset: (page - 1) * DEFAULT_LIMIT,
		begTimestamp: date.from.toISOString(),
		endTimestamp: date.to.toISOString(),
		feedbackType: search.type,
		feedbackTags: search.tags,
		feedbackMeta: search.meta
	};
	const [searchResponse, countResponse] = await Promise.all([
		api().Feedback.search(organization.id, options),
		api().Feedback.count(organization.id, options)
	]);
	return {
		feedbacks: {
			count: countResponse.total,
			pages: Math.ceil(countResponse.total / DEFAULT_LIMIT),
			items: searchResponse.feedbacks
		}
	};
};
