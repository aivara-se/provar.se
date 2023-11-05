import { OrganizationRepository } from '$lib/server';
import type { Session } from '@auth/core/types';
import createSlug from 'slug';
import type { Actions } from './$types';
import { error, redirect } from '@sveltejs/kit';

async function assertOrganizationMembership(organizationId: string, userId: string) {
	const organization = await OrganizationRepository.findById(organizationId);
	if (!organization) {
		throw error(403);
	}
	const isMember = organization?.members.includes(userId);
	if (!isMember) {
		throw error(403);
	}
}

export const actions: Actions = {
	updateOrganization: async ({ locals, request }) => {
		const session = (await locals.getSession()) as Session;
		const data = await request.formData();
		const orgId = data.get('id') as string;
		const name = data.get('name') as string;
		const prod = data.get('prod') === 'true';
		await assertOrganizationMembership(orgId, session.user.id);
		const slug = createSlug(name);
		await OrganizationRepository.update(orgId, { name, slug, prod });
		throw redirect(303, `/org/${slug}/preferences`);
	}
};
