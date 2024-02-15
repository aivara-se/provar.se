import { OrganizationService } from '$lib/server/organization';
import type { Session } from '@auth/core/types';
import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
	default: async ({ locals, request }) => {
		const session = (await locals.auth()) as Session;
		const data = await request.formData();
		const name = data.get('name') as string;
		const orgId = await OrganizationService.create(session.user.id, name);
		redirect(303, `/org/${orgId}`);
	}
};
