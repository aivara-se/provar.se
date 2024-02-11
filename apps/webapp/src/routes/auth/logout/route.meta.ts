import type { Route } from '$lib/types';

const route: Route = {
	getName: () => 'Logout',
	getPath: () => '/auth/logout'
};

export default route;
