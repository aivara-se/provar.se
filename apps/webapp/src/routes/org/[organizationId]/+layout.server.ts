import type { Session } from '@auth/core/types';
import type { LayoutServerLoad } from './$types';
import { OrganizationRepository } from '$lib/server';

export const load: LayoutServerLoad = async (event) => {
	const session = (await event.locals.getSession()) as Session;
	const organizations = await OrganizationRepository.findByMember(session.user.id);
	return { organizations };
};
