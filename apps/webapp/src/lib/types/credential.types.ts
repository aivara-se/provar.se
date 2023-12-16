/**
 * Credentials are used when collecting feedback from users.
 */
export interface Credential {
	id: string;
	name: string;
	organizationId: string;
	key: string;
	createdAt: Date;
}
