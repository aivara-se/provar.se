/**
 * The type of feedback. This determines what data is stored in the database,
 * hot it is analyzed and how it is displayed on UI.
 */
export enum FeedbackType {
	Text = 'text',
	Rate = 'rate'
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
	projectId?: string;
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
 * Data stored in the database for rate feedback.
 */
export interface RateFeedbackData {
	rate: number;
}

/**
 * A feedback type that stores a single text field.
 */
export interface RateFeedback extends BaseFeedback {
	type: FeedbackType.Rate;
	data: RateFeedbackData;
}

/**
 * A union type of all available feedback types.
 */
export type Feedback = TextFeedback | RateFeedback;
