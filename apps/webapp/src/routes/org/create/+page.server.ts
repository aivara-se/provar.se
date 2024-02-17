import { OrganizationService } from '$lib/server/organization';
import { error, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import type { Actions, PageServerLoad } from './$types';
import { schema } from './(forms)/CreateOrganizationForm.schema';

export const load: PageServerLoad = async () => {
	const form = await superValidate(zod(schema));
	return { CreateOrganizationForm: form };
};

export const actions: Actions = {
	createOrganization: async (event) => {
		const session = await event.locals.auth();
		if (!session || !session.user?.id) {
			return error(403);
		}
		const form = await superValidate(event.request, zod(schema));
		const { name } = form.data;
		const orgId = await OrganizationService.create(session.user.id, name);
		redirect(302, `/org/${orgId}`);
	}
};
