import Papa from 'papaparse';

/**
 * Parse a CSV string into an array of objects.
 */
export function parse(csv: string): Record<string, string>[] {
	const { data } = Papa.parse<Record<string, string>>(csv, { header: true });
	return data;
}
