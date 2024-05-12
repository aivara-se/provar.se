import { defineRoute } from '$lib/client/routes';
import { Users2Icon } from 'lucide-svelte';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-settings-members',
	tabs: ['org-settings-organization', 'org-settings-members', 'org-settings-credentials'],
	icon: Users2Icon,
	parent: parentRoute,
	getName: () => 'Members',
	getPath: (params) => `/org/${params.slug}/settings/members`
});
