/**
 * Invitations to join organizations.
 */
export interface Invitation {
	id: string;
	key: string;
	name: string;
	email: string;
	organizationId: string;
	createdAt: Date;
}
