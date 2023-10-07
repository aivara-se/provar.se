import {
	AUTH_GITHUB_ID,
	AUTH_GITHUB_SECRET,
	AUTH_GOOGLE_ID,
	AUTH_GOOGLE_SECRET,
	AUTH_SECRET
} from '$env/static/private';
import { getMongoConnection } from '$lib/server/mongo';
import GitHub from '@auth/core/providers/github';
import Google from '@auth/core/providers/google';
import { MongoDBAdapter } from '@auth/mongodb-adapter';
import { SvelteKitAuth } from '@auth/sveltekit';
import { redirect, type Handle } from '@sveltejs/kit';
import { sequence } from '@sveltejs/kit/hooks';

const authorization: Handle = async ({ event, resolve }) => {
	if (!event.url.pathname.startsWith('/auth/')) {
		const session = await event.locals.getSession();
		if (!session) {
			throw redirect(303, '/auth/login');
		}
	}
	return resolve(event);
};

const sveltekitauth = SvelteKitAuth({
	adapter: MongoDBAdapter(getMongoConnection()),
	providers: [
		Google({
			clientId: AUTH_GOOGLE_ID,
			clientSecret: AUTH_GOOGLE_SECRET,
			allowDangerousEmailAccountLinking: true
		}),
		GitHub({
			clientId: AUTH_GITHUB_ID,
			clientSecret: AUTH_GITHUB_SECRET,
			allowDangerousEmailAccountLinking: true
		})
	],
	secret: AUTH_SECRET,
	session: {
		strategy: 'jwt'
	},
	pages: {
		signIn: '/auth/login',
		signOut: '/auth/logout',
		error: '/auth/error',
		verifyRequest: '/auth/verify',
		newUser: '/'
	},
	trustHost: true
});

export const handle = sequence(sveltekitauth, authorization);
