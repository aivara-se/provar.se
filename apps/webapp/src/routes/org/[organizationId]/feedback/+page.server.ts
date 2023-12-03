import {
	CredentialService,
	FeedbackRepository,
	FeedbackService,
	ProvarAPIService,
	getSelectedOrganization
} from '$lib/server';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const { organization } = await event.parent();
	const feedbacks = await FeedbackRepository.findByOrganization(organization.id);
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
		await ProvarAPIService.importFeedback(credential.key, signedUrl);
	}
};
