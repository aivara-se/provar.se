import { getSelectedOrganization } from '$lib/server/action-utils';
import { CredentialService } from '$lib/server/credential';
import { FeedbackService } from '$lib/server/feedback';
import { importFeedback } from '$lib/server/provar-api';
import type { Actions } from './$types';

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
