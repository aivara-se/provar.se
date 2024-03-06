import { CSAT_DATA } from './csat.data.js';

export type CSATQuestionType = 'rating-11p';

export interface CSATFeedback {
	'question-type': CSATQuestionType;
	'response-data': number;
	'response-text'?: string;
}

export function generateCSATFeedback(): CSATFeedback {
	const responseData = Math.floor(Math.random() * 11);
	const result: CSATFeedback = {
		'question-type': 'rating-11p',
		'response-data': responseData
	};
	const shouldHaveText = Math.random() > 0.25;
	if (!shouldHaveText) {
		return result;
	}
	const offset = CSAT_DATA.length * (0.25 - Math.random() * 0.5);
	const indexWithOffset = Math.floor((responseData / 10) * CSAT_DATA.length + offset);
	const index = Math.max(0, Math.min(CSAT_DATA.length - 1, indexWithOffset));
	result['response-text'] = CSAT_DATA[index];
	return result;
}
