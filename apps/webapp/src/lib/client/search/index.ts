import { browser } from '$app/environment';
import { goto } from '$app/navigation';
import uniq from 'lodash/uniq';
import { get, readonly, writable, type Readable } from 'svelte/store';
import type { FeedbackType } from '../types';

/**
 * The name of the search query parameter in the URL.
 */
export const SEARCH_QUERY_PARAM = 'search';

/**
 * Parsed search query
 */
export interface SearchQuery {
	type: FeedbackType[];
	text: string[];
	tags: Record<string, string>;
	meta: Record<string, string>;
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
		if (part === 'cnps') {
			types.push('cnps');
		} else if (part === 'csat') {
			types.push('csat');
		} else if (part === 'text') {
			types.push('text');
		}
	}
	return types;
}

/**
 * Get the initial value of the search store from the current URL.
 */
function getInitialValue(): SearchQuery {
	if (!browser) {
		return parseSearch('');
	}
	const value = new URLSearchParams(window.location.search).get('search') || '';
	return parseSearch(value);
}

/**
 * Create a store that represents the current search query parameters.
 */
const store = writable<SearchQuery>(getInitialValue());

/**
 * Get an empty search query object.
 */
function getEmptyValue(): SearchQuery {
	return { type: [], text: [], tags: {}, meta: {} };
}

/**
 * Get the current search query parameters from the store.
 */
function getValue(): SearchQuery {
	return get(store);
}

/**
 * Set the current search query parameters on the store and update the URL.
 */
function setValue(value: SearchQuery) {
	store.set(value);
	updateUrl(value);
}

/**
 * Update the URL with the given search query parameters.
 */
function updateUrl(value: SearchQuery) {
	if (!browser) {
		return;
	}
	const str = stringifySearch(value);
	const url = new URL(window.location.href);
	if (str) {
		url.searchParams.set(SEARCH_QUERY_PARAM, str);
	} else {
		url.searchParams.delete(SEARCH_QUERY_PARAM);
	}
	goto(url.toString(), { replaceState: true, keepFocus: true });
}

/**
 * Get a read only version of a store with the search query parameters.
 */
export function getSearchStore(): Readable<SearchQuery> {
	return readonly(store);
}

/**
 * Get the current search query parameters from the store as an object.
 */
export function getSearchValue(): SearchQuery {
	return getValue();
}

/**
 * Get the current search query parameters from the store as a string.
 */
export function getSearchString(): string {
	return stringifySearch(getValue());
}

/**
 * Replaces the search query parameter in the store.
 */
export function setSearchValue(...values: Partial<SearchQuery>[]) {
	let acc = getEmptyValue();
	for (const value of values) {
		acc = mergeSearchQuery(acc, value);
	}
	setValue(acc);
}

/**
 * Replaces the search query parameter in the store.
 */
export function setSearchString(...values: string[]) {
	let acc = getEmptyValue();
	for (const value of values) {
		acc = mergeSearchQuery(acc, parseSearch(value));
	}
	setValue(acc);
}

/**
 * Merges the search query parameter in the store with the new value.
 */
export function mergeSearchValue(...values: Partial<SearchQuery>[]) {
	let acc = getValue();
	for (const value of values) {
		acc = mergeSearchQuery(acc, value);
	}
	setValue(acc);
}

/**
 * Merges the search query parameter and navigates to the explore page.
 */
export function exploreWithValue(organizationId: string, ...values: Partial<SearchQuery>[]) {
	mergeSearchValue(...values);
	const url = new URL(window.location.href);
	url.pathname = `/org/${organizationId}/explore`;
	goto(url.toString());
}
