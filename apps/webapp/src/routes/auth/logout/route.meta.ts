import { defineRoute } from '$lib/client/routes';
import parentRoute from '../../route.meta';

export default defineRoute({
	id: 'auth-logout',
	hidden: true,
	parent: parentRoute,
	getName: () => 'Logout',
	getPath: () => '/auth/logout'
});
