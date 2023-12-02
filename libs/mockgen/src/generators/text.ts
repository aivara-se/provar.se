import { TEXT_DATA } from './text.data.js';

export interface TextFeedback {
	text: string;
}

export function generateTextFeedback(): TextFeedback {
	const result: TextFeedback = {
		text: TEXT_DATA[Math.floor(Math.random() * TEXT_DATA.length)]
	};
	return result;
}
