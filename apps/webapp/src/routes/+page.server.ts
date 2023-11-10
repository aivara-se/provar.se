import { OrganizationRepository } from '$lib/server';
import type { Session } from '@auth/core/types';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const session = (await event.locals.getSession()) as Session;
	const organizations = await OrganizationRepository.findByMember(session.user.id);
	if (organizations.length === 0) {
		throw redirect(302, `/org/create`);
	}
	if (organizations.length === 1) {
		throw redirect(302, `/org/${organizations[0].id}`);
	}
	return { organizations };
};
