import { PUBLIC_PROVAR_APP_URL } from '$env/static/public';

/**
 * Creates a link that can be used to accept an invitation.
 */
export function createInvitationLink(key: string): string {
	return `${PUBLIC_PROVAR_APP_URL}/auth/accept/${key}`;
}
