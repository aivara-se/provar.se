import { defineRoute } from '$lib/client/routes';
import { Building2Icon } from 'lucide-svelte';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-settings-organization',
	tabs: ['org-settings-organization', 'org-settings-members', 'org-settings-credentials'],
	icon: Building2Icon,
	parent: parentRoute,
	getName: () => 'Organization',
	getPath: (params) => `/org/${params.slug}/settings/organization`
});
