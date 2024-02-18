import { defineRoute } from '$lib/client/routes';
import { SettingsIcon } from 'lucide-svelte';

export default defineRoute({
	id: 'org-settings',
	icon: SettingsIcon,
	sidebar: { weight: 30, mobile: true },
	getName: () => 'Settings',
	getPath: (params) => `/org/${params.organizationId}/settings`
});
