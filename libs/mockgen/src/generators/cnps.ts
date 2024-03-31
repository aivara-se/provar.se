import { CNPS_DATA } from './cnps.data.js';

export type CNPSQuestionType = 'rating-11p';

export interface CNPSFeedback {
	'question-type': CNPSQuestionType;
	'response-data': number;
	'response-text': string;
}

export function generateCNPSFeedback(): CNPSFeedback {
	const responseData = Math.floor(Math.random() * 11);
	const result: CNPSFeedback = {
		'question-type': 'rating-11p',
		'response-data': responseData,
		'response-text': ''
	};
	const shouldHaveText = Math.random() > 0.25;
	if (!shouldHaveText) {
		return result;
	}
	const offset = CNPS_DATA.length * (0.25 - Math.random() * 0.5);
	const indexWithOffset = Math.floor((responseData / 10) * CNPS_DATA.length + offset);
	const index = Math.max(0, Math.min(CNPS_DATA.length - 1, indexWithOffset));
	result['response-text'] = CNPS_DATA[index];
	return result;
}
