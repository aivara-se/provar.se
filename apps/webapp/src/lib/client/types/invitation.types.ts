/**
 * Invitations to join organizations.
 */
export interface Invitation {
	id: string;
	organizationId: string;
	createdAt: string;
	createdBy: string;
	expiresAt: string;
	acceptedAt?: string;
	secret: string;
	name: string;
	email: string;
	avatar: string;
}
