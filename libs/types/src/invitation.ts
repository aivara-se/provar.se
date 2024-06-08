export interface Invitation {
	id: string;
	organization_id: string;
	name: string;
	email: string;
	avatar: string;
	secret: string;
	accepted: boolean;
	createdAt: string;
	updatedAt: string;
	expiresAt: string;
}
