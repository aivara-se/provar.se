import { parseDateRange } from '$lib/client/dates';
import { parseSearch } from '$lib/client/search';
import type { PageLoad } from './$types';

const DEFAULT_LIMIT = 10;

const EMPTY_RESPONSE = {
	count: 0,
	items: [],
	pages: 0
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
		return EMPTY_RESPONSE;
	}
	const date = parseDateRange(range);
	const page = Number.parseInt(event.url.searchParams.get('page') ?? '1');
	const search = parseSearch(event.url.searchParams.get('search') ?? '');
	const options = { date, page, limit: DEFAULT_LIMIT, search };
	console.log('TODO: load feedbacks with options', organization, options);
	const count = 0;
	return {
		feedbacks: {
			count: count,
			pages: Math.ceil(count / DEFAULT_LIMIT),
			items: []
		}
	};
};
