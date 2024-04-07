/**
 * Result of validating a CSV file data.
 */
export interface Result {
	total: number;
	valid: number;
}

/**
 * Validate the data from a CSV file.
 */
export function validate(data: Record<string, string>[]): Result {
	let valid = 0;
	for (const row of data) {
		if (!isValidType(row.type)) {
			continue;
		}
		// TODO: add more validation rules
		valid++;
	}
	return { total: data.length, valid };
}

/**
 * Check if the type field is valid.
 */
function isValidType(type: string): boolean {
	return ['text', 'cnps', 'csat'].includes(type);
}
