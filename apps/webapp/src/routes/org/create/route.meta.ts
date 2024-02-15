import { defineRoute } from '$lib/client/routes';

export default defineRoute({
	id: 'org-create',
	getName: () => 'Create organization',
	getPath: () => '/org/create'
});
