import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-list',
	hidden: true,
	parent: parentRoute,
	getName: () => 'Organizations',
	getPath: () => '/org'
});
