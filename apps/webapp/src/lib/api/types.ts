export interface Organization {
	id: string;
	name: string;
	environment: 'staging' | 'production';
}

export interface Feedback {
	id: string;
	organizationId: string;
	createdAt: number;
	projectId: string | null;
	features: { [key: string]: boolean };
	metadata: { [key: string]: unknown };
	formdata: { [key: string]: unknown };
}
