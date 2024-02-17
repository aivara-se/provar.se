import { env } from '$env/dynamic/private';
import { getMongoClient } from '$lib/server/database';
import { EmailProvider } from '$lib/server/email-utils';
import GitHubProvider from '@auth/core/providers/github';
import GoogleProvider from '@auth/core/providers/google';
import { MongoDBAdapter } from '@auth/mongodb-adapter';
import { SvelteKitAuth } from '@auth/sveltekit';

export const { handle, signIn, signOut } = SvelteKitAuth({
	adapter: MongoDBAdapter(getMongoClient()),
	providers: [
		EmailProvider(),
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
});
