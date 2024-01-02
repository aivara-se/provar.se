import { FeedbackType } from '$lib/types';
import uniq from 'lodash/uniq';

/**
 * Parsed search query
 */
export interface SearchQuery {
	type: FeedbackType[];
	text: string[];
	tags: Record<string, string>;
	meta: Record<string, unknown>;
}

/**
 * Merges two search queries together using the first param as the base.
 */
export function mergeSearchQuery(x: SearchQuery, y: Partial<SearchQuery>): SearchQuery {
	return {
		type: y.type ? uniq([...x.type, ...y.type]) : x.type,
		text: y.text ? [...x.text, ...y.text] : x.text,
		tags: y.tags ? { ...x.tags, ...y.tags } : x.tags,
		meta: y.meta ? { ...x.meta, ...y.meta } : x.meta
	};
}

/**
 * Removes leading and trailing quotes from the given string.
 */
function dequote(str: string) {
	const match = str.match(/^(['"])(.*)\1$/);
	if (match && match[2]) {
		return match[2];
	}
	return str;
}

/**
 * Converts a search string to a search query.
 */
export function parseSearch(search: string): SearchQuery {
	const query: SearchQuery = { type: [], text: [], tags: {}, meta: {} };
	const parts = search.match(/(?:[^\s"]+|"[^"]*")+/g) || [search];
	parts.forEach((part) => {
		const segments = part.split('=');
		const key = segments[0].trim();
		const val = segments.slice(1).join('=')?.trim();
		const maybeType = parseTypes(dequote(val));
		if (key.toLowerCase() === 'type' && maybeType.length) {
			query.type.push(...maybeType);
		} else if (key[0] === '#' && val) {
			query.tags[dequote(key.slice(1))] = dequote(val);
		} else if (key[0] === '$' && val) {
			query.meta[dequote(key.slice(1))] = dequote(val);
		} else if (val) {
			query.meta[dequote(key)] = dequote(val);
		} else {
			query.text.push(dequote(key));
		}
	});
	return query;
}

/**
 * Converts a search query to a search string.
 */
export function stringifySearch(query: SearchQuery) {
	const parts = [];
	if (query.type.length) {
		parts.push(`type=${query.type.join(',')}`);
	}
	for (const [key, value] of Object.entries(query.tags)) {
		parts.push(`#${key}="${value}"`);
	}
	for (const [key, value] of Object.entries(query.meta)) {
		parts.push(`$${key}="${value}"`);
	}
	for (const text of query.text) {
		parts.push(text);
	}
	return parts.join(' ');
}

/**
 * Identifies the type of feedback from the given value.
 */
function parseTypes(typeStr: string): FeedbackType[] {
	if (!typeStr) {
		return [];
	}
	const types: FeedbackType[] = [];
	const parts = typeStr.split(',').map((v) => v.trim().toLowerCase());
	for (const part of parts) {
		if (part === 'cnps' || part === 'nps') {
			types.push(FeedbackType.CNPS);
		} else if (part === 'csat') {
			types.push(FeedbackType.CSAT);
		} else if (part === 'text' || part === 'txt') {
			types.push(FeedbackType.Text);
		}
	}
	return types;
}
