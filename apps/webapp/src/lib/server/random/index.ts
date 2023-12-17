/**
 * Generates a secure random string.
 */
export function secureRandom(count: number): string {
	const array = new Uint8Array(count);
	crypto.getRandomValues(array);
	return Buffer.from(array).toString('base64url');
}
