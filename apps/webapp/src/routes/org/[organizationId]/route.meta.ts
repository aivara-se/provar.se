import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-root',
	hidden: true,
	parent: parentRoute,
	getName: () => 'Organization',
	getPath: (params) => `/org/${params.organizationId}`
});
