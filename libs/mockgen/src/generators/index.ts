import { generateCNPSFeedback } from './cnps.js';
import { generateCSATFeedback } from './csat.js';
import { generateRandomDate } from './date.js';
import { generateMetadata } from './meta.js';
import { generateTextFeedback } from './text.js';
import { generateUserdata } from './user.js';

export type FeedbackType = 'text' | 'csat' | 'cnps';

const FeedbackGeneratorsMap = {
	text: generateTextFeedback,
	csat: generateCSATFeedback,
	cnps: generateCNPSFeedback
};

export function generateFeedback(type: FeedbackType, period: [Date, Date]) {
	const [start, end] = period;
	return {
		type,
		time: generateRandomDate(start, end),
		data: FeedbackGeneratorsMap[type](),
		meta: Math.random() > 0.2 ? generateMetadata() : {},
		user: Math.random() > 0.2 ? generateUserdata() : {}
	};
}
