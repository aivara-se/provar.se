import { env } from '$env/dynamic/private';
import { getMongoClient } from '$lib/server/database';
import { EmailService, VerifyLoginTemplate } from '$lib/server/email-utils';
import EmailProvider from '@auth/core/providers/email';
import GitHubProvider from '@auth/core/providers/github';
import GoogleProvider from '@auth/core/providers/google';
import { MongoDBAdapter } from '@auth/mongodb-adapter';
import { SvelteKitAuth } from '@auth/sveltekit';
import { redirect, type Handle } from '@sveltejs/kit';
import { sequence } from '@sveltejs/kit/hooks';
import process from 'process';

process.on('SIGINT', function () {
	process.exit();
});

process.on('SIGTERM', function () {
	process.exit();
});

const authorization: Handle = async ({ event, resolve }) => {
	// Note: needed for Cloudflare Pages 404.html
	if (event.url.pathname === '/[fallback]') {
		return resolve(event);
	}
	if (event.url.pathname.startsWith('/auth/')) {
		return resolve(event);
	}
	const session = await event.locals.getSession();
	if (!session) {
		redirect(303, '/auth/login');
	}
	return resolve(event);
};

const sveltekitauth: Handle = (input) =>
	SvelteKitAuth({
		adapter: MongoDBAdapter(getMongoClient()),
		providers: [
			EmailProvider({
				async sendVerificationRequest(params) {
					await EmailService.send({
						toEmail: params.identifier,
						options: { link: params.url },
						template: VerifyLoginTemplate
					});
				}
				// eslint-disable-next-line @typescript-eslint/no-explicit-any
			}) as any,
			GoogleProvider({
				clientId: env.AUTH_GOOGLE_ID,
				clientSecret: env.AUTH_GOOGLE_SECRET,
				allowDangerousEmailAccountLinking: true
			}),
			GitHubProvider({
				clientId: env.AUTH_GITHUB_ID,
				clientSecret: env.AUTH_GITHUB_SECRET,
				allowDangerousEmailAccountLinking: true
			})
		],
		secret: env.AUTH_SECRET,
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
		trustHost: true,
		callbacks: {
			async session(params) {
				if (params.session.user && 'token' in params) {
					params.session.user.id = params.token.sub as string;
				}
				return params.session;
			}
		}
	})(input);

export const handle = sequence(sveltekitauth, authorization);
