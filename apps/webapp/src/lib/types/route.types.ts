export interface Route {
	// The parent route
	parent?: Route;

	// The page name with given route parameters
	getName: (params: Record<string, string>) => string;

	// The page path with given route parameters
	getPath: (params: Record<string, string>) => string;
}
