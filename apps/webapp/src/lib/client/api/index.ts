import { PUBLIC_PROVAR_API_URL } from '$env/static/public';
import { DefaultFetcher, ProvarClient, _overrideConfig } from '@provar/provar-js';
import { accessToken, getAccessToken } from '../auth/token.store';

if (PUBLIC_PROVAR_API_URL) {
	_overrideConfig({ baseUrl: PUBLIC_PROVAR_API_URL });
}

/**
 * This variable is used to cache the Provar client.
 */
let snapshot = createClient();

/**
 * This store is used to keep the Provar client in sync with the access token.
 */
accessToken.subscribe(() => {
	snapshot = createClient();
});

/**
 * This function is used to get the current snapshot of the Provar client.
 */
export function api() {
	return snapshot;
}

/**
 * This function is used to create a new Provar client.
 */
export function createClient() {
	const token = getAccessToken() || '';
	return new ProvarClient({ token });
}
