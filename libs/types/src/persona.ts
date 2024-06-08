export interface Persona {
	id: string;
	organization_id: string;
	name: string;
	email: string;
	avatar: string;
	metadata: Record<string, string>;
	userdata: Record<string, string>;
	createdAt: string;
	updatedAt: string;
}
