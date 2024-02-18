import { defineRoute } from '$lib/client/routes';
import { SearchIcon } from 'lucide-svelte';

export default defineRoute({
	id: 'org-explore',
	icon: SearchIcon,
	sidebar: { weight: 10, mobile: true },
	getName: () => 'Explore',
	getPath: (params) => `/org/${params.organizationId}/explore`,
	isActive: (path, params) => path.startsWith(`/org/${params.organizationId}/explore`)
});
