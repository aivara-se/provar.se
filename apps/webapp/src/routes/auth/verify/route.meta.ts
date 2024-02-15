import { defineRoute } from '$lib/client/routes';

export default defineRoute({
	id: 'auth-verify',
	getName: () => 'Verify',
	getPath: () => '/auth/verify'
});
