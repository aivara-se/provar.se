import { defineRoute } from '$lib/client/routes';

export default defineRoute({
	id: 'org-selector',
	getName: () => 'Organizations',
	getPath: () => '/org'
});
