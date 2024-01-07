import { getSelectedOrganization } from '$lib/server/action-utils';
import { CredentialService } from '$lib/server/credential';
import { FeedbackRepository, FeedbackService } from '$lib/server/feedback';
import { importFeedback } from '$lib/server/provar-api';
import { parseDateRange } from '$lib/shared/dates';
import { parseSearch } from '$lib/shared/search';
import type { Actions, PageServerLoad } from './$types';

const DEFAULT_LIMIT = 10;

/**
 * Loads feedbacks to display on the page.
 */
export const load: PageServerLoad = async (event) => {
	const { organization } = await event.parent();
	const page = Number.parseInt(event.url.searchParams.get('page') ?? '1');
	const date = parseDateRange({ from: event.params.dateFrom, to: event.params.dateTo });
	const search = parseSearch(event.url.searchParams.get('search') ?? '');
	const options: FeedbackRepository.FindOptions = { date, page, limit: DEFAULT_LIMIT, search };
	const [items, count, summary] = await Promise.all([
		FeedbackRepository.findByOrganization(organization.id, options),
		FeedbackRepository.countByOrganization(organization.id, options),
		FeedbackService.getFeedbackSummary(organization.id, options)
	]);
	return {
		feedbacks: {
			count: count,
			items: items,
			summary: summary,
			pages: Math.ceil(count / DEFAULT_LIMIT)
		}
	};
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
