import { defineRoute } from '$lib/client/routes';
import { FoldersIcon } from 'lucide-svelte';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-projects',
	icon: FoldersIcon,
	parent: parentRoute,
	sidebar: { weight: 20, mobile: true },
	getName: () => 'Projects',
	getPath: (params) => `/org/${params.organizationId}/projects`
});
