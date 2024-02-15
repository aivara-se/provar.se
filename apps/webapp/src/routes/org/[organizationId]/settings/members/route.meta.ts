import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-settings-members',
	parent: parentRoute,
	getName: () => 'Members',
	getPath: (params) => `/org/${params.organizationId}/settings/members`,
	isActive: (path, params) => path.startsWith(`/org/${params.organizationId}/settings/members`)
});
