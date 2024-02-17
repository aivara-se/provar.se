import { sha256 } from 'js-sha256';

/**
 * Creates a link that can be used to get the gravatar image
 */
export function createGravatarLink(email: string): string {
	const hash = sha256(
		email
			.trim()
			.toLowerCase()
			.replace(/\+[^@]+@/, '@')
	);
	return `https://www.gravatar.com/avatar/${hash}?s=64&d=mp`;
}
