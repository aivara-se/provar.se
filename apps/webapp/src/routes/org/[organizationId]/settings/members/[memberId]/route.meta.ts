import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'member-details',
	tabs: ['org-settings-organization', 'org-settings-members', 'org-settings-credentials'],
	parent: parentRoute,
	getName: () => 'Member Details',
	getPath: (params) => `/org/${params.organizationId}/settings/members/${params.memberId}`,
	isActive: (path, params) =>
		path.startsWith(`/org/${params.organizationId}/settings/members/${params.memberId}`)
});
