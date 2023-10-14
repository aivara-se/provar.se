import { OrganizationRepository, type Environment } from '$lib/server';
import type { Session } from '@auth/core/types';
import { redirect } from '@sveltejs/kit';
import createSlug from 'slug';
import type { Actions } from './$types';

export const actions = {
	default: async ({ locals, request }) => {
		const session = (await locals.getSession()) as Session;
		const data = await request.formData();

		const name = data.get('name') as string;
		const slug = createSlug(name);
		const environment = data.get('environment') as Environment;

		await OrganizationRepository.create(session.user.id, {
			name,
			slug,
			environment
		});

		throw redirect(303, `/org/${slug}`);
	}
} satisfies Actions;
