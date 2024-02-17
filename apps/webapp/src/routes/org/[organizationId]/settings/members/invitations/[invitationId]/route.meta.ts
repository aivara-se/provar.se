import { defineRoute } from '$lib/client/routes';
import parentRoute from '../../route.meta';

export default defineRoute({
	id: 'invitation-details',
	tabs: ['org-settings-organization', 'org-settings-members', 'org-settings-credentials'],
	parent: parentRoute,
	getName: () => 'Invitation Details',
	getPath: (params) =>
		`/org/${params.organizationId}/settings/members/invitations/${params.invitationId}`,
	isActive: (path, params) =>
		path.startsWith(
			`/org/${params.organizationId}/settings/members/invitations/${params.invitationId}`
		)
});
