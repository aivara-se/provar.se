import { defineRoute } from '$lib/client/routes';
import { SearchIcon } from 'lucide-svelte';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-explore',
	icon: SearchIcon,
	parent: parentRoute,
	sidebar: { weight: 10, mobile: true },
	getName: () => 'Explore',
	getPath: (params) => `/org/${params.organizationId}/explore`
});
