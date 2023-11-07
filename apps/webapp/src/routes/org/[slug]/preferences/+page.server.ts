import { OrganizationRepository, CredentialRepository, type Organization } from '$lib/server';
import type { Session } from '@auth/core/types';
import createSlug from 'slug';
import type { Actions, PageServerLoad, RequestEvent } from './$types';
import { error, redirect } from '@sveltejs/kit';

/**
 * Loads data required for all tabs on organization preferences page.
 */
export const load: PageServerLoad = async ({ locals, params }) => {
	const session = (await locals.getSession()) as Session;
	const organization = await OrganizationRepository.findBySlug(params.slug);
	if (!organization || isMember(organization, session.user.id)) {
		throw error(403);
	}
	const credentials = await CredentialRepository.findByOrganization(organization.id);
	return { credentials: credentials };
};

/**
 * Asserts that the user is a member of the organization.
 */
function isMember(organization: Organization, userId: string): boolean {
	return !organization?.members.includes(userId);
}

/**
 * Decorator to ensure that the user is a member of the organization.
 */
function withOrganization(fn: (organization: Organization, event: RequestEvent) => unknown) {
	return async (event: RequestEvent) => {
		const session = (await event.locals.getSession()) as Session;
		const organization = await OrganizationRepository.findBySlug(event.params.slug);
		if (!organization || isMember(organization, session.user.id)) {
			throw error(403);
		}
		return fn(organization, event);
	};
}

export const actions: Actions = {
	/**
	 * Updates the organization's name and other details.
	 */
	updateOrganization: withOrganization(async (organization, { request }) => {
		const data = await request.formData();
		const name = String(data.get('name'));
		const prod = data.get('prod') === 'true';
		const slug = createSlug(name);
		await OrganizationRepository.update(organization.id, { name, slug, prod });
		throw redirect(303, `/org/${slug}/preferences`);
	}),
	/**
	 * Creates a new client credential for the organization.
	 */
	createCredential: withOrganization(async (organization, { request }) => {
		const data = await request.formData();
		await CredentialRepository.create({
			name: String(data.get('name')),
			organizationId: organization.id,
			clientId: secureRandom(64),
			clientSecret: secureRandom(64)
		});
	}),
	/**
	 * Creates a new client credential for the organization.
	 */
	revokeCredential: withOrganization(async (organization, { request }) => {
		const data = await request.formData();
		const id = String(data.get('id'));
		await CredentialRepository.revoke(organization.id, id);
	})
};

/**
 * Generates a secure random string.
 */
function secureRandom(count: number): string {
	const array = new Uint8Array(count);
	crypto.getRandomValues(array);
	return Buffer.from(array).toString('base64url');
}
