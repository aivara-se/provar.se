import { OrganizationService } from '$lib/server/organization';
import type { Session } from '@auth/core/types';
import { redirect } from '@sveltejs/kit';
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
		const session = (await event.locals.auth()) as Session;
		const form = await superValidate(event.request, zod(schema));
		const { name } = form.data;
		const orgId = await OrganizationService.create(session.user.id, name);
		redirect(303, `/org/${orgId}`);
	}
};
