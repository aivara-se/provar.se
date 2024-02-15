import { defineRoute } from '$lib/client/routes';
import { SettingsIcon } from 'lucide-svelte';

export default defineRoute({
	id: 'org-settings',
	sidebar: { weight: 30, icon: SettingsIcon, mobile: true },
	getName: () => 'Settings',
	getPath: (params) => `/org/${params.organizationId}/settings`,
	isActive: (path, params) => path.startsWith(`/org/${params.organizationId}/settings`)
});
