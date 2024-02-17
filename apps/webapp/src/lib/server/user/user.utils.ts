/**
 * Get the name of a user from their name or email.
 */
export function getUserName(name: string, email: string): string {
	if (name) {
		return name;
	}
	return email.split('@')[0];
}

/**
 * Get the initials of a user from their name or email.
 */
export function getUserInitials(name: string, email: string): string {
	if (name) {
		return name
			.split(' ')
			.map((name) => name[0])
			.join('')
			.slice(0, 2)
			.toLocaleUpperCase();
	}
	return email.slice(0, 2).toLocaleUpperCase();
}
