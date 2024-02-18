import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-onboard',
	parent: parentRoute,
	getName: () => 'Onboarding',
	getPath: (params) => `/org/${params.organizationId}/onboard`
});
