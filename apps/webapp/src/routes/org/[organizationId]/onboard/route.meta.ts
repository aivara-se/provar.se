import { defineRoute } from '$lib/client/routes';

export default defineRoute({
	id: 'org-onboard',
	getName: () => 'Onboarding',
	getPath: (params) => `/org/${params.organizationId}/onboard`
});
