import { defineRoute } from '$lib/client/routes';

export default defineRoute({
	id: 'auth-logout',
	getName: () => 'Logout',
	getPath: () => '/auth/logout'
});
