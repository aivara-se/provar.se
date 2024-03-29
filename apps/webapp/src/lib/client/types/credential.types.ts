/**
 * Credentials are used when collecting feedback from users.
 */
export interface Credential {
	id: string;
	organizationId: string;
	createdAt: string;
	createdBy: string;
	modifiedAt: string;
	lastUsedAt?: string;
	name: string;
	secret: string;
}
