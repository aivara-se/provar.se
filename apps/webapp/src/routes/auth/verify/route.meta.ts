import type { Route } from '$lib/types';

const route: Route = {
	getName: () => 'Verify',
	getPath: () => '/auth/verify'
};

export default route;
