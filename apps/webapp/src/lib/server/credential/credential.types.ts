/**
 * Credentials are used when collecting feedback from users.
 */
export interface Credential {
	id: string;
	name: string;
	organizationId: string;
	clientId: string;
	clientSecret: string;
}
