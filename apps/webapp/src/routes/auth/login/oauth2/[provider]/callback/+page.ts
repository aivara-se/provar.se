import { browser } from '$app/environment';
import { redirect } from '@sveltejs/kit';
import { isValidOAuth2State, verifyOAuth2Login } from '$lib/client/auth';
import type { PageLoad } from './$types';

export const load: PageLoad = async (event) => {
	if (!browser) {
		return;
	}
	if (!isValidOAuth2State) {
		redirect(302, `/auth/login/failed`);
	}
	const state = event.url.searchParams.get('state');
	const code = event.url.searchParams.get('code');
	if (!state || !code) {
		redirect(302, `/auth/login/failed`);
	}
	await verifyOAuth2Login(event.params.provider, state, code);
	redirect(302, `/org`);
};
