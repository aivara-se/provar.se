import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'org-create',
	hidden: true,
	parent: parentRoute,
	getName: () => 'Create organization',
	getPath: () => '/org/create'
});
