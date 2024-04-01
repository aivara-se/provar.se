import { TEXT_DATA } from './text.data.js';

export type TextQuestionType = 'default';

export interface TextFeedback {
	'question-type': TextQuestionType;
	'response-text': string;
}

export function generateTextFeedback(): TextFeedback {
	const result: TextFeedback = {
		'question-type': 'default',
		'response-text': TEXT_DATA[Math.floor(Math.random() * TEXT_DATA.length)]
	};
	return result;
}
