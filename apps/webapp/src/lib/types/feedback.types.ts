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
 * Feedback tags are used to add additional metadata to feedback messages.
 */
export type FeedbackTags = Record<string, string>;

/**
 * The base feedback type other feedback types extend from. This type is not exported
 * and should not be used directly. Use the union type `Feedback` instead.
 */
interface BaseFeedback {
	id: string;
	organizationId: string;
	projectId?: string;
	createdAt: number;
	type: FeedbackType;
	tags?: FeedbackTags;
	data: unknown;
}

/**
 * Data stored in the database for text feedback.
 */
export interface TextFeedbackData {
	text: string;
}

/**
 * A feedback type that stores a single text field.
 */
export interface TextFeedback extends BaseFeedback {
	type: FeedbackType.Text;
	data: TextFeedbackData;
}

/**
 * Data stored in the database for cNPS feedback.
 */
export interface CNPSFeedbackData {
	cnps: number;
}

/**
 * A feedback type that stores a single cNPS value.
 */
export interface CNPSFeedback extends BaseFeedback {
	type: FeedbackType.CNPS;
	data: CNPSFeedbackData;
}

/**
 * Data stored in the database for CSAT feedback.
 */
export interface CSATFeedbackData {
	csat: number;
}

/**
 * A feedback type that stores a single CSAT value.
 */
export interface CSATFeedback extends BaseFeedback {
	type: FeedbackType.CSAT;
	data: CSATFeedbackData;
}

/**
 * A union type of all available feedback types.
 */
export type Feedback = TextFeedback | CNPSFeedback | CSATFeedback;
