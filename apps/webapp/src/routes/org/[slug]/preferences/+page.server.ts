import { OrganizationRepository, type Organization } from '$lib/server';
import type { Session } from '@auth/core/types';
import createSlug from 'slug';
import type { Actions } from './$types';
import { error, redirect } from '@sveltejs/kit';

async function assertOrganizationMembership(organization: Organization, userId: string) {
	const isMember = organization?.members.includes(userId);
	if (!isMember) {
		throw error(403);
	}
}

export const actions: Actions = {
	updateOrganization: async ({ locals, request, params }) => {
		const session = (await locals.getSession()) as Session;
		const data = await request.formData();
		const name = data.get('name') as string;
		const prod = data.get('prod') === 'true';
		const organization = await OrganizationRepository.findBySlug(params.slug);
		if (!organization) {
			throw error(403);
		}
		await assertOrganizationMembership(organization, session.user.id);
		const slug = createSlug(name);
		await OrganizationRepository.update(organization.id, { name, slug, prod });
		throw redirect(303, `/org/${slug}/preferences`);
	}
};
