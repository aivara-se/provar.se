/**
 * The user/s work on project/s
 */
export interface Project {
	id: string;
	createdAt: Date;
	name: string;
	description: string;
	organizationId: string;
	collectionGoal: number | null;
}
