import { generateCNPSFeedback } from './generators/cnps.js';
import { generateCSATFeedback } from './generators/csat.js';
import { generateRandomDate } from './generators/date.js';
import { generateMetadata } from './generators/meta.js';
import { generateRandomTags } from './generators/tags.js';
import { generateTextFeedback } from './generators/text.js';
import { generateUserdata } from './generators/user.js';

/**
 * Export helper function to generate the default range
 */
export { getDefaultPeriod } from './generators/time.js';

/**
 * Feedback types
 */
export type FeedbackType = 'text' | 'csat' | 'cnps';

/**
 * Map of feedback types => feedback data generators
 */
const FeedbackGeneratorsMap = {
	text: generateTextFeedback,
	csat: generateCSATFeedback,
	cnps: generateCNPSFeedback
};

/**
 * Generates a feedback object
 */
export function generateFeedback(type: FeedbackType, period: [Date, Date]) {
	const [start, end] = period;
	return {
		type,
		time: generateRandomDate(start, end),
		data: FeedbackGeneratorsMap[type](),
		meta: Math.random() > 0.2 ? generateMetadata() : {},
		user: Math.random() > 0.2 ? generateUserdata() : {},
		tags: Math.random() > 0.2 ? generateRandomTags() : {}
	};
}
