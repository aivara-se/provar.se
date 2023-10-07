import { SvelteKitAuth } from '@auth/sveltekit';
import GitHub from '@auth/core/providers/github';
import { AUTH_GITHUB_ID, AUTH_GITHUB_SECRET, AUTH_SECRET } from '$env/static/private';
import { redirect, type Handle } from '@sveltejs/kit';
import { sequence } from '@sveltejs/kit/hooks';

const authorization: Handle = async ({ event, resolve }) => {
	if (!event.url.pathname.startsWith('/auth')) {
		const session = await event.locals.getSession();
		if (!session) {
			throw redirect(303, '/auth/login');
		}
	}
	return resolve(event);
};

const sveltekitauth = SvelteKitAuth({
	providers: [GitHub({ clientId: AUTH_GITHUB_ID, clientSecret: AUTH_GITHUB_SECRET })],
	secret: AUTH_SECRET,
	session: {
		strategy: 'jwt'
	},
	pages: {
		signIn: '/auth/login',
		signOut: '/auth/logout',
		error: '/auth/error',
		verifyRequest: '/auth/verify',
		newUser: '/auth/new-user'
	},
	trustHost: true
});

export const handle = sequence(sveltekitauth, authorization);
