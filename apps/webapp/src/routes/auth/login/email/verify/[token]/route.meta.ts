import { defineRoute } from '$lib/client/routes';
import parentRoute from '../../route.meta';

export default defineRoute({
	id: 'auth-login-email-verify',
	hidden: true,
	parent: parentRoute,
	getName: () => 'Login with email verification',
	getPath: () => '/auth/login/email/verify/:key'
});
