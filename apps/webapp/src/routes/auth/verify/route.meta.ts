import { defineRoute } from '$lib/client/routes';
import parentRoute from '../../route.meta';

export default defineRoute({
	id: 'auth-verify',
	hidden: true,
	parent: parentRoute,
	getName: () => 'Verify',
	getPath: () => '/auth/verify'
});
