import { defineRoute } from '$lib/client/routes';
import parentRoute from '../route.meta';

export default defineRoute({
	id: 'project-details',
	parent: parentRoute,
	getName: () => 'Project Details',
	getPath: (params) => `/org/${params.organizationId}/projects/${params.projectId}`,
	isActive: (path, params) =>
		path.startsWith(`/org/${params.organizationId}/projects/${params.projectId}`)
});
