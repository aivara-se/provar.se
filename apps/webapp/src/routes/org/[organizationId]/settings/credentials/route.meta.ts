import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-settings-credentials',
	parent: parentRoute,
	getName: () => 'Credentials',
	getPath: (params) => `/org/${params.organizationId}/settings/credentials`,
	isActive: (path, params) => path.startsWith(`/org/${params.organizationId}/settings/credentials`)
});
