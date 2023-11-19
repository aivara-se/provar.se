/**
 * The user/s work on project/s
 */
export interface Project {
	id: string;
	name: string;
	description: string;
	organizationId: string;
	collectionGoal: number | null;
}
