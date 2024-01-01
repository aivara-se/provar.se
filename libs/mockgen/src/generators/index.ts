import { generateCNPSFeedback } from './cnps.js';
import { generateCSATFeedback } from './csat.js';
import { generateMetadata } from './meta.js';
import { generateTextFeedback } from './text.js';
import { generateUserdata } from './user.js';

export type FeedbackType = 'text' | 'csat' | 'cnps';

const FeedbackGeneratorsMap = {
	text: generateTextFeedback,
	csat: generateCSATFeedback,
	cnps: generateCNPSFeedback
};

function randomDate(start: Date, end: Date): Date {
	const diff = end.getTime() - start.getTime();
	const newDiff = diff * Math.random();
	return new Date(start.getTime() + newDiff);
}

export function generateFeedback(type: FeedbackType, period: [Date, Date]) {
	const [start, end] = period;
	return {
		type,
		time: randomDate(start, end),
		data: FeedbackGeneratorsMap[type](),
		meta: generateMetadata(),
		user: generateUserdata()
	};
}
