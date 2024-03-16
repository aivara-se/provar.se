import { PUBLIC_PROVAR_API_URL } from '$env/static/public';
import { ProvarClient, _overrideConfig } from '@provar/provar-js';
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
	const token = getAccessToken();
	snapshot = createClient(token ?? '');
});

/**
 * This function is used to get the current snapshot of the Provar client.
 */
export function getApi() {
	return snapshot;
}

/**
 * This function is used to create a new Provar client.
 */
function createClient(token = '') {
	return new ProvarClient({ token });
}
