import { getSelectedOrganization } from '$lib/server/action-utils';
import { CredentialService } from '$lib/server/credential';
import { FeedbackRepository, FeedbackService } from '$lib/server/feedback';
import { importFeedback } from '$lib/server/provar-api';
import { parseSearch } from '$lib/shared/search';
import type { Actions, PageServerLoad } from './$types';

const ITEMS_PER_PAGE = 10;
const DEFAULT_RANGE = '30d';

/**
 * Converts a range string to a date range.
 */
function parseRange(range: string): { from: Date; to: Date } {
	const now = new Date();
	const val = Number.parseInt(range.slice(0, -1));
	const mul = 24 * 60 * 60 * 1000;
	return { from: new Date(now.getTime() - val * mul), to: now };
}

/**
 * Loads feedbacks to display on the page.
 */
export const load: PageServerLoad = async (event) => {
	const { organization } = await event.parent();
	const page = Number.parseInt(event.url.searchParams.get('page') ?? '1');
	const date = parseRange(event.url.searchParams.get('range') ?? DEFAULT_RANGE);
	const search = parseSearch(event.url.searchParams.get('search') ?? '');
	const options: FeedbackRepository.FindOptions = { date, page, limit: ITEMS_PER_PAGE, search };
	const { items, count } = await FeedbackRepository.findByOrganization(organization.id, options);
	const feedbacks = { count, pages: Math.ceil(count / ITEMS_PER_PAGE), items };
	return { feedbacks };
};

export const actions: Actions = {
	/**
	 * Create a signed URL for uploading a CSV file with feedbacks.
	 */
	createImportUrl: async (event) => {
		const organization = await getSelectedOrganization(event);
		const { signedUrl, fileName } = await FeedbackService.createUploadUrl(organization.id);
		return { signedUrl, fileName };
	},

	/**
	 * Start processing the SCV file uploaded to Google Cloud Storage.
	 */
	processImportFile: async (event) => {
		const organization = await getSelectedOrganization(event);
		const data = await event.request.formData();
		const name = String(data.get('name'));
		const { signedUrl } = await FeedbackService.createReadableUrl(organization.id, name);
		const credential = await CredentialService.getImportCredential(organization.id);
		await importFeedback(credential.key, signedUrl);
	}
};
