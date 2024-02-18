import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-feedback-details',
	parent: parentRoute,
	getName: () => 'Feedback Details',
	getPath: (params) => `/org/${params.organizationId}/explore/${params.feedbackId}`
});
