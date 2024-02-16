/**
 * Generate a random string id
 */
export function createId(length = 10) {
	let result = '';
	while (result.length < length) {
		result += Math.random().toString(16).slice(2);
	}
	return result.slice(0, length);
}
