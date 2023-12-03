import { PROVAR_API_URL } from '$env/static/private';

/**
 * Sends user feedback to Provar API to be imported.
 */
export async function importFeedback(apiKey: string, link: string): Promise<void> {
	const url = `${PROVAR_API_URL}/feedback/import`;
	const fetchResponse = await fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${apiKey}`
		},
		body: JSON.stringify({ link })
	});
	if (fetchResponse.status >= 400) {
		throw new Error(`Provar: request failed with status ${fetchResponse.status}`);
	}
}
