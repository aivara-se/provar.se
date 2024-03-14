import { clearAccessToken, getAccessToken } from './token.store';

/**
 * This function is used to sign in with magic email link.
 */
export function signInWithEmail(email: string) {
	// TODO: implement
	console.log('email:', email);
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
	// TODO: Redirect to the sign in page
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
