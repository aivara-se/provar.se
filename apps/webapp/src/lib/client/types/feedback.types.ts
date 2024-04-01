/**
 * The type of feedback. This determines what data is stored in the database,
 * hot it is analyzed and how it is displayed on UI.
 */
export enum FeedbackType {
	Text = 'text',
	CNPS = 'cnps',
	CSAT = 'csat'
}

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
	image?: string;
}

/**
 * Feedback tags are used to group feedback messages into categories
 */
export type FeedbackTags = Record<string, string>;

/**
 * The base feedback type other feedback types extend from. This type is not exported
 * and should not be used directly. Use the union type `Feedback` instead.
 */
interface BaseFeedback {
	id: string;
	organizationId: string;
	createdAt: string;
	questionType: string;
	feedbackTime: string;
	feedbackType: FeedbackType;
	feedbackData: Record<string, string>;
	feedbackTags: Record<string, string>;
	feedbackMeta: Record<string, string>;
	feedbackUser: Record<string, string>;
}

/**
 * Data stored in the database for text feedback.
 */
export interface TextFeedbackData {
	'question-type': 'default';
	'response-text': string;
	[key: string]: string;
}

/**
 * A feedback type that stores a single text field.
 */
export interface TextFeedback extends BaseFeedback {
	feedbackType: FeedbackType.Text;
	feedbackData: TextFeedbackData;
}

/**
 * Data stored in the database for cNPS feedback.
 */
export interface CNPSFeedbackData {
	'question-type': 'rating-11p';
	'response-text': string;
	'response-data': string;
	[key: string]: string;
}

/**
 * A feedback type that stores a single cNPS value.
 */
export interface CNPSFeedback extends BaseFeedback {
	feedbackType: FeedbackType.CNPS;
	feedbackData: CNPSFeedbackData;
}

/**
 * Data stored in the database for CSAT feedback.
 */
export interface CSATFeedbackData {
	'question-type': 'rating-11p';
	'response-text': string;
	'response-data': string;
	[key: string]: string;
}

/**
 * A feedback type that stores a single CSAT value.
 */
export interface CSATFeedback extends BaseFeedback {
	feedbackType: FeedbackType.CSAT;
	feedbackData: CSATFeedbackData;
}

/**
 * A union type of all available feedback types.
 */
export type Feedback = TextFeedback | CNPSFeedback | CSATFeedback;
