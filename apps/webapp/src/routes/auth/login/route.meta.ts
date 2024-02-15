import { defineRoute } from '$lib/client/routes';

export default defineRoute({
	id: 'auth-login',
	getName: () => 'Login',
	getPath: () => '/auth/login'
});
