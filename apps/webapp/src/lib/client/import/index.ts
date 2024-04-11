import Papa from 'papaparse';
import type { FeedbackType } from '../types';

/**
 * The data structure of a row in the feedback CSV.
 */
export interface FeedbackRow {
	type: FeedbackType;
	time: string;
	[key: `data.${string}`]: string;
	[key: `meta.${string}`]: string;
	[key: `user.${string}`]: string;
	[key: `tags.${string}`]: string;
}

/**
 * The data structure of the parsed feedback data.
 */
export interface FeedbackItem {
	type: FeedbackType;
	time: string;
	data: Record<string, string>;
	tags: Record<string, string>;
	meta: Record<string, string>;
	user: Record<string, string>;
}

/**
 * Parse a CSV string into an array of FeedbackItem objects.
 */
export function parseCsv(csv: string): FeedbackItem[] {
	const { data } = Papa.parse<FeedbackRow>(csv, { header: true });
	const items: FeedbackItem[] = data.map((row) => ({
		type: row.type,
		time: row.time,
		data: mapPrefix('data.', row),
		tags: mapPrefix('tags.', row),
		meta: mapPrefix('meta.', row),
		user: mapPrefix('user.', row)
	}));
	return items;
}

/**
 * Result of validating a CSV file data.
 */
export interface ValidateResult {
	total: number;
	valid: number;
}

/**
 * Validate the data from a CSV file.
 */
export function validateCsv(data: FeedbackItem[]): ValidateResult {
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
function isValidType(type: string): type is FeedbackType {
	return ['text', 'cnps', 'csat'].includes(type);
}

/**
 * Converts a flattened set of keys into a nested object.
 */
function mapPrefix(prefix: string, obj: FeedbackRow): Record<string, string> {
	const result: Record<string, string> = {};
	for (const key in obj) {
		if (key.startsWith(prefix)) {
			result[key.replace(prefix, '')] = obj[key as keyof FeedbackRow];
		}
	}
	return result;
}
