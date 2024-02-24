/**
 * Invitations to join organizations.
 */
export interface Invitation {
	id: string;
	organizationId: string;
	createdAt: Date;
	acceptedAt?: Date;
	key: string;
	name: string;
	link: string;
	email: string;
	image: string;
}
