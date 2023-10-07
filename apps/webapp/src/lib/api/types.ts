export interface User {
	id: string;
	name: string;
}

export type UserRole = 'owner' | 'admin' | 'member';

export interface Membership {
	id: string;
	role: UserRole;
}

export type Environment = 'staging' | 'production';

export interface Organization {
	id: string;
	slug: string;
	name: string;
	environment: Environment;
	members: Membership[];
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
