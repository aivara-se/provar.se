import { defineRoute } from '$lib/client/routes';
import { SettingsIcon } from 'lucide-svelte';

export default defineRoute({
	id: 'org-settings',
	icon: SettingsIcon,
	tabs: ['org-settings-organization', 'org-settings-members', 'org-settings-credentials'],
	sidebar: { weight: 30, mobile: true },
	getName: () => 'Settings',
	getPath: (params) => `/org/${params.organizationId}/settings`,
	isActive: (path, params) => path.startsWith(`/org/${params.organizationId}/settings`)
});