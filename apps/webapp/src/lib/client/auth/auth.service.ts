import { goto } from '$app/navigation';
import { getApi } from '../api';
import { clearAccessToken, getAccessToken, setAccessToken } from './token.store';

/**
 * This function is used to sign in with magic email link.
 */
export async function signInWithEmail(email: string) {
	await getApi().EmailAuthentication.prepare({ email });
	goto('/auth/login/email/sent');
}

/**
 * This function is used to verify the email login token.
 */
export async function verifyEmailLogin(token: string) {
	const response = await getApi().EmailAuthentication.confirm({ token });
	setAccessToken(response.accessToken);
	goto('/');
}

/**
 * This function is used to sign in with Google OAuth provider.
 */
export function signInWithGoogle() {
	// TODO: implement
}

/**
 * This function is used to sign in with Github OAuth provider.
 */
export function signInWithGithub() {
	// TODO: implement
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
	const response = await getApi().Authentication.whoami();
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
