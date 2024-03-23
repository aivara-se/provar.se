import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-import',
	parent: parentRoute,
	getName: () => 'Import',
	getPath: (params) => `/org/${params.slug}/import`
});
