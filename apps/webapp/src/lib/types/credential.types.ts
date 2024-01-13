/**
 * Credentials are used when collecting feedback from users.
 */
export interface Credential {
	id: string;
	createdAt: Date;
	lastUsedAt?: Date;
	name: string;
	organizationId: string;
	key: string;
}
