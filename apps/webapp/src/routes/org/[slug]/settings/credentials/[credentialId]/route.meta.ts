import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'credential-details',
	// tabs: ['org-settings-organization', 'org-settings-members', 'org-settings-credentials'],
	tabs: ['org-settings-organization', 'org-settings-credentials'],
	parent: parentRoute,
	getName: () => 'Credential Details',
	getPath: (params) => `/org/${params.slug}/settings/credentials/${params.credentialId}`
});
