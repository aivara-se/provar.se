/**
 * Invitations to join organizations.
 */
export interface Invitation {
	id: string;
	key: string;
	email: string;
	organizationId: string;
}
