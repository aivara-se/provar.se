import { CSAT_DATA } from './csat.data.js';

export interface CSATFeedback {
	csat: number;
	text?: string;
}

export function generateCSATFeedback(): CSATFeedback {
	const result: CSATFeedback = {
		csat: Math.random()
	};
	const shouldHaveText = Math.random() > 0.5;
	if (!shouldHaveText) {
		return result;
	}
	const offset = CSAT_DATA.length * (0.25 - Math.random() * 0.5);
	const indexWithOffset = Math.floor(result.csat * CSAT_DATA.length + offset);
	const index = Math.max(0, Math.min(CSAT_DATA.length - 1, indexWithOffset));
	result.text = CSAT_DATA[index];
	return result;
}
