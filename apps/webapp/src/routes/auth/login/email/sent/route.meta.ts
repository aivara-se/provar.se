import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'auth-login-email-sent',
	hidden: true,
	parent: parentRoute,
	getName: () => 'Login email sent message',
	getPath: () => '/auth/login/email/sent'
});
