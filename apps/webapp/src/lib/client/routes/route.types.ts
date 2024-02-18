import type { Action } from '$lib/client/action';

export interface Route {
	id: string;

	// NOTE: The `SvelteComponent` type is not supported by icon libraries
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	icon?: any;

	// The parent route
	parent?: Route;

	// Show child routes as tabs on the route page
	tabs?: string[];

	// Flag: the route should be shown on the sidebar
	sidebar?: { weight: number; mobile?: boolean };

	// Flag: the route should be hidden from breadcrumbs
	hidden?: boolean;

	// The page name with given route parameters
	getName: (params: Record<string, string>) => string;

	// The page path with given route parameters
	getPath: (params: Record<string, string>) => string;

	// Checks if the route is active for the given path and parameters
	isActive?: (path: string, params: Record<string, string>) => boolean;

	// Get actions the user can perform on a route and parent routes
	getActions?(path: string, params: Record<string, string>): Action[];
}
