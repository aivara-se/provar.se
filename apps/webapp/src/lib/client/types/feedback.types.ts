/**
 * The type of feedback. This determines what data is stored in the database,
 * hot it is analyzed and how it is displayed on UI.
 */
export type FeedbackType = 'text' | 'cnps' | 'csat';

/**
 * Feedback metadata are additional metadata attached to feedback messages
 */
export type FeedbackMeta = {
	[key: string]: unknown;
};

/**
 * Feedback user is used to identify the user that sent the feedback.
 */
export interface FeedbackUser {
	id?: string;
	name?: string;
	email?: string;
	avatar?: string;
}

/**
 * Feedback tags are used to group feedback messages into categories
 */
export type FeedbackTags = Record<string, string>;

/**
 * The base feedback type other feedback types extend from. This type is not exported
 * and should not be used directly. Use the union type `Feedback` instead.
 */
export interface Feedback {
	id: string;
	organizationId: string;
	createdAt: string;
	feedbackTime: string;
	feedbackType: FeedbackType;
	feedbackData: Record<string, string>;
	feedbackTags: Record<string, string>;
	feedbackMeta: Record<string, string>;
	feedbackUser: Record<string, string>;
}
