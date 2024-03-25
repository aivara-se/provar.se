import { defineRoute } from '$lib/client/routes';
import { SettingsIcon } from 'lucide-svelte';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-settings',
	icon: SettingsIcon,
	parent: parentRoute,
	sidebar: { weight: 30, mobile: true },
	getName: () => 'Settings',
	getPath: (params) => `/org/${params.slug}/settings`
});
