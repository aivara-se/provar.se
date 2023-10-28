import { ProjectRepository } from '$lib/server';
import type { Actions } from './$types';
import { redirect } from '@sveltejs/kit';

export const actions: Actions = {
	default: async ({ request }) => {
		const data = await request.formData();
		const name = data.get('name') as string;
		const organizationId = data.get('organizationId') as string;
		// TODO: make sure user belongs to the organization
		await ProjectRepository.create({ name, organizationId });
		// TODO: redirect user to the created project
		throw redirect(303, `/`);
	}
};
