import { defineRoute } from '$lib/client/routes';
import parentRoute from '../../route.meta';

export default defineRoute({
	id: 'auth-login',
	hidden: true,
	parent: parentRoute,
	getName: () => 'Login',
	getPath: () => '/auth/login'
});
