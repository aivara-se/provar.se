import type { FeedbackType } from './feedback.types';

export enum ProjectStatus {
	Backlog = 'backlog',
	Collecting = 'collecting',
	Completed = 'completed'
}

/**
 * The user/s work on project/s
 */
export interface Project {
	id: string;
	organizationId: string;
	createdAt: Date;
	name: string;
	description: string;
	status: ProjectStatus;
	feedbackType: FeedbackType;
}
