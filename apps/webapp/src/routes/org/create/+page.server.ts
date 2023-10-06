import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions = {
	default: async (event) => {
		// TODO: Check if user is logged in
		// console.log('cookies', event.cookies.getAll());
		// console.log('session', await event.locals.getSession());
		throw redirect(302, `/org/org-a`);
	}
} satisfies Actions;
