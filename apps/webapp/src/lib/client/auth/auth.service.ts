import { goto } from '$app/navigation';
import { PUBLIC_PROVAR_API_URL } from '$env/static/public';
import { api } from '../api';
import { clearAccessToken, getAccessToken, setAccessToken } from './token.store';

/**
 * This local storage key is used to store the access token
 */
const STATE_PARAM_KEY = 'provar:oauth2_state';

/**
 * For how long the state parameter is valid.
 */
const STATE_PARAM_TTL = 1000 * 60 * 60; // 1 hour

/**
 * This function is used to sign in with magic email link.
 */
export async function signInWithEmail(email: string) {
	await api().EmailAuthentication.prepare({ email });
	goto('/auth/login/email/sent');
}

/**
 * This function is used to verify the email login token.
 */
export async function verifyEmailLogin(token: string) {
	const response = await api().EmailAuthentication.confirm({ token });
	setAccessToken(response.accessToken);
	goto('/');
}

/**
 * This function is used to sign in with Google OAuth provider.
 */
export function signInWithGoogle() {
	const state = setupOAuth2State();
	location.href = `${PUBLIC_PROVAR_API_URL}/auth/oauth2/google/login?state=${state}`;
}

/**
 * This function is used to sign in with Github OAuth provider.
 */
export function signInWithGithub() {
	const state = setupOAuth2State();
	location.href = `${PUBLIC_PROVAR_API_URL}/auth/oauth2/github/login?state=${state}`;
}

/**
 * This function is used to clear the access token from local storage
 */
export function signOut() {
	clearAccessToken();
	goto('/auth/login');
}

/**
 * This function is used to get the current user details.
 */
export async function getCurrentUser() {
	const token = getAccessToken();
	if (!token) {
		return null;
	}
	const response = await api().Authentication.whoami();
	if (response.type !== 'user') {
		return null;
	}
	return response.user;
}

/**
 * This function is used to check whether the user is authorized to access
 * the current page. If the user is not authorized, it will return false.
 */
export function canAccessRoute(pathname: string, whitelist: string[]): boolean {
	const token = getAccessToken();
	if (token) {
		return true;
	}
	const whitelisted = whitelist.some((path) => pathname.startsWith(path));
	if (whitelisted) {
		return true;
	}
	return false;
}

/**
 * This function is used to create a new OAuth2 state and store it in local storage.
 */
function setupOAuth2State() {
	let val = '';
	while (val.length < 32) {
		val += Math.random().toString(16).substring(7);
	}
	val = val.substring(0, 32);
	const obj = JSON.stringify({ val, ts: Date.now() });
	localStorage.setItem(STATE_PARAM_KEY, obj);
	return val;
}

/**
 * This function is used to validate the OAuth2 state parameter in local storage.
 * And exchange OAuth2 login params for an access token using the API.
 */
export function isValidOAuth2State(provider: string, state: string) {
	const obj = localStorage.getItem(STATE_PARAM_KEY);
	if (!obj) {
		return false;
	}
	localStorage.removeItem(STATE_PARAM_KEY);
	const { val, ts } = JSON.parse(obj);
	if (val !== state) {
		return false;
	}
	if (Date.now() - ts > STATE_PARAM_TTL) {
		return false;
	}
	return true;
}

/**
 * This function is used to complete the oauth2 login flow.
 */
export async function verifyOAuth2Login(provider: string, state: string, code: string) {
	const response = await api().OAuth2Authentication.confirm(provider, { state, code });
	setAccessToken(response.accessToken);
	goto('/');
}
