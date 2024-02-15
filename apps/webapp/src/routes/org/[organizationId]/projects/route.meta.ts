import { defineRoute } from '$lib/client/routes';
import { FoldersIcon } from 'lucide-svelte';

export default defineRoute({
	id: 'org-projects',
	icon: FoldersIcon,
	sidebar: { weight: 20, mobile: true },
	getName: () => 'Projects',
	getPath: (params) => `/org/${params.organizationId}/projects`,
	isActive: (path, params) => path.startsWith(`/org/${params.organizationId}/projects`)
});
