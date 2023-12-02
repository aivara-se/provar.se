import { CNPS_DATA } from './cnps.data.js';

export interface CNPSFeedback {
	cnps: number;
	text?: string;
}

export function generateCNPSFeedback(): CNPSFeedback {
	const result: CNPSFeedback = {
		cnps: Math.random()
	};
	const shouldHaveText = Math.random() > 0.5;
	if (!shouldHaveText) {
		return result;
	}
	const offset = CNPS_DATA.length * (0.25 - Math.random() * 0.5);
	const indexWithOffset = Math.floor(result.cnps * CNPS_DATA.length + offset);
	const index = Math.max(0, Math.min(CNPS_DATA.length - 1, indexWithOffset));
	result.text = CNPS_DATA[index];
	return result;
}
