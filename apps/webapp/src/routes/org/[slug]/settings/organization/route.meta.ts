import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-settings-organization',
	// tabs: ['org-settings-organization', 'org-settings-members', 'org-settings-credentials'],
	tabs: ['org-settings-organization'],
	parent: parentRoute,
	getName: () => 'Organization',
	getPath: (params) => `/org/${params.slug}/settings/organization`
});
