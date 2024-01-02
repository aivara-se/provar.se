import { browser } from '$app/environment';
import { goto } from '$app/navigation';
import {
	mergeSearchQuery,
	parseSearch,
	stringifySearch,
	type SearchQuery
} from '$lib/shared/search';
import { get, readonly, writable, type Readable } from 'svelte/store';

/**
 * The name of the search query parameter in the URL.
 */
export const SEARCH_QUERY_PARAM = 'search';

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
	const url = new URL(window.location.toString());
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
