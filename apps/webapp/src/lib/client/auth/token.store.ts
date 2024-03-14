import { writable } from 'svelte/store';

/**
 * This local storage key is used to store the access token
 */
const ACCESS_TOKEN_KEY = 'provar:access_token';

/**
 * This interface represents the decoded access token
 */
export interface AccessTokenClaims {
	sub: string;
	exp: number;
}

/**
 * This store is used to keep the access token in sync with local storage
 */
export const accessToken = writable(getAccessTokenClaims());

/**
 * This function is used to parse the access token from a string.
 */
export function parseAccessToken(tokenStr: string): AccessTokenClaims | null {
	try {
		const token = JSON.parse(atob(tokenStr.split('.')[1]));
		if (Date.now() >= token.exp * 1000) {
			return null;
		}
		return token;
	} catch (e) {
		return null;
	}
}

/**
 * This function is used to get the access token from local storage.
 * Returns null if the access token is not available. Also returns
 * null if the access token is not valid or expired.
 */
export function getAccessToken(): string | null {
	const tokenStr = localStorage.getItem(ACCESS_TOKEN_KEY);
	if (!tokenStr) {
		return null;
	}
	const token = parseAccessToken(tokenStr);
	if (!token) {
		return null;
	}
	return tokenStr;
}

/**
 * This function is used to get the access token from local storage
 * and decode it. Returns null if the access token is not available.
 * Also returns null if the access token is not valid or expired.
 */
export function getAccessTokenClaims(): AccessTokenClaims | null {
	const tokenStr = localStorage.getItem(ACCESS_TOKEN_KEY);
	if (!tokenStr) {
		return null;
	}
	return parseAccessToken(tokenStr);
}

/**
 * This function is used to set the access token in local storage.
 */
export function setAccessToken(tokenStr: string) {
	const token = parseAccessToken(tokenStr);
	if (!token) {
		return clearAccessToken();
	}
	localStorage.setItem(ACCESS_TOKEN_KEY, tokenStr);
	accessToken.set(token);
}

/**
 * This function is used to clear the access token from local storage.
 */
export function clearAccessToken() {
	localStorage.removeItem(ACCESS_TOKEN_KEY);
	accessToken.set(null);
}
