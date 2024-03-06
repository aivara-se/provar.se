import { TEXT_DATA } from './text.data.js';

export interface TextFeedback {
	'response-text': string;
}

export function generateTextFeedback(): TextFeedback {
	const result: TextFeedback = {
		'response-text': TEXT_DATA[Math.floor(Math.random() * TEXT_DATA.length)]
	};
	return result;
}
