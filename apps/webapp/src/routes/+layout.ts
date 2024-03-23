import { canAccessRoute, getCurrentUser } from '$lib/client/auth';
import { redirect } from '@sveltejs/kit';
import { api } from '../lib/client/api';
import type { LayoutLoad } from './$types';

/**
 * Disable server-side rendering.
 */
export const ssr = false;

/**
 * A list of routes that do not require authentication.
 */
const AUTH_WHITELIST = ['/auth'];

/**
 * Redirect users to the login page if they are not authenticated.
 */
export const load: LayoutLoad = async (event) => {
	if (!canAccessRoute(event.url.pathname, AUTH_WHITELIST)) {
		redirect(302, '/auth/login');
	}
	const user = await getCurrentUser();
	if (!user) {
		return { user: null, organizations: [] };
	}
	const organizations = await api().Organization.list();
	return { user, organizations };
};
