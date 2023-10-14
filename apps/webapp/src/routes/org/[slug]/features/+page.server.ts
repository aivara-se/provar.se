import { OrganizationRepository } from '$lib/server';
import type { Session } from '@auth/core/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const session = (await event.locals.getSession()) as Session;
	const organizations = await OrganizationRepository.findByMember(session.user.id);
	return { organizations };
};
