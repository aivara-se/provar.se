model Response {
	id: string;
	organization_id: string;
	persona_id: string | null;
	project_id: string;
	timestamp: string;
	metadata: Record<string>;
	userdata: Record<string>;
	createdAt: string;
	updatedAt: string;
}
