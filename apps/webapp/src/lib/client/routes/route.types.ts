export interface Route {
	id: string;

	// The parent route
	parent?: Route;

	// Flag: the route should be shown on the sidebar
	// NOTE: The `SvelteComponent` type is not supported by icon libraries
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	sidebar?: { weight: number; icon: any; mobile?: boolean };

	// The page name with given route parameters
	getName: (params: Record<string, string>) => string;

	// The page path with given route parameters
	getPath: (params: Record<string, string>) => string;

	// Checks if the route is active for the given path and parameters
	isActive?: (path: string, params: Record<string, string>) => boolean;
}
