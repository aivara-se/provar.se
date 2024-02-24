/**
 * Credentials are used when collecting feedback from users.
 */
export interface Credential {
	id: string;
	organizationId: string;
	createdAt: Date;
	lastUsedAt?: Date;
	name: string;
	key: string;
}
