import { defineRoute } from '$lib/client/routes';
import { KeyRoundIcon } from 'lucide-svelte';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-settings-credentials',
	tabs: ['org-settings-organization', 'org-settings-members', 'org-settings-credentials'],
	icon: KeyRoundIcon,
	parent: parentRoute,
	getName: () => 'Credentials',
	getPath: (params) => `/org/${params.slug}/settings/credentials`
});
