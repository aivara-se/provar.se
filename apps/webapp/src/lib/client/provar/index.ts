import { PUBLIC_PROVAR_API_KEY } from '$env/static/public';
import { ProvarClient } from '@provar/provar-js';

/**
 * Provar client instance.
 */
const provarClient = new ProvarClient({ apiKey: PUBLIC_PROVAR_API_KEY });

/**
 * Sends user feedback to Provar.
 */
export async function sendTextFeedback(text: string) {
	await provarClient.sendText(text);
}
