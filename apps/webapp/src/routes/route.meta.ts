import { signOut } from '@auth/sveltekit/client';
import { defineRoute } from '$lib/client/routes';
import type { FrontentAction } from '$lib/client/action';

const LogoutAction: FrontentAction = {
	id: 'logout',
	type: 'frontend',
	name: 'Logout',
	exec: () => signOut()
};

export default defineRoute({
	id: 'root',
	hidden: true,
	getName: () => 'Root',
	getPath: () => '/',
	getActions: () => [LogoutAction]
});
