import { redirect, type Handle } from '@sveltejs/kit';
import { sequence } from '@sveltejs/kit/hooks';
import process from 'process';
import { handle as sveltekitauth } from './auth';

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

export const handle = sequence(sveltekitauth, authorization);
