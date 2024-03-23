import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-export',
	parent: parentRoute,
	getName: () => 'Export',
	getPath: (params) => `/org/${params.slug}/export`
});
