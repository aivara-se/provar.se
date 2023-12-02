import { CNPS_DATA } from './cnps.data.js';
import { CSAT_DATA } from './csat.data.js';

export const TEXT_DATA: string[] = [];

/**
 * Populate TEXT_DATA with a mix of CNPS_DATA and CSAT_DATA in a way that
 * the sorting of the resulting array is as close as possible to the sorting
 * of the two source arrays.
 */
(function () {
	const length = CNPS_DATA.length + CSAT_DATA.length;
	const sources = [CNPS_DATA, CSAT_DATA];
	const indexes = sources.map(() => 0);

	// mutates `indexes`
	const next = () => {
		let selectedScore = indexes[0] / sources[0].length;
		let selectedValue = sources[0][indexes[0]];
		let selectedIndex = 0;
		for (let i = 1; i < sources.length; i++) {
			const currentScore = indexes[i] / sources[i].length;
			if (currentScore < selectedScore) {
				selectedScore = currentScore;
				selectedValue = sources[i][indexes[i]];
				selectedIndex = i;
			}
		}
		indexes[selectedIndex]++;
		return selectedValue;
	};

	for (let i = 0; i < length; i++) {
		TEXT_DATA[i] = next();
	}
})();
