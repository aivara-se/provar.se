import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-settings-members',
	tabs: ['org-settings-organization', 'org-settings-members', 'org-settings-credentials'],
	parent: parentRoute,
	getName: () => 'Members',
	getPath: (params) => `/org/${params.organizationId}/settings/members`
});
