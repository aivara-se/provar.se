import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'credential-details',
	tabs: ['org-settings-organization', 'org-settings-members', 'org-settings-credentials'],
	parent: parentRoute,
	getName: () => 'Details',
	getPath: (params) => `/org/${params.organizationId}/settings/credentials/${params.credentialId}`,
	isActive: (path, params) =>
		path.startsWith(`/org/${params.organizationId}/settings/credentials/${params.credentialId}`)
});
