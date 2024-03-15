import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'auth-login-email',
	hidden: true,
	parent: parentRoute,
	getName: () => 'Login with email',
	getPath: () => '/auth/login/email'
});
