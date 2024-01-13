/**
 * Invitations to join organizations.
 */
export interface Invitation {
	id: string;
	createdAt: Date;
	acceptedAt?: Date;
	key: string;
	name: string;
	link: string;
	email: string;
	organizationId: string;
}
