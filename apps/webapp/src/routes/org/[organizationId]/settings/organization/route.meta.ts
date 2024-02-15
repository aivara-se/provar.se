import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-settings-organization',
	parent: parentRoute,
	getName: () => 'Organization',
	getPath: (params) => `/org/${params.organizationId}/settings/organization`,
	isActive: (path, params) => path.startsWith(`/org/${params.organizationId}/settings/organization`)
});
