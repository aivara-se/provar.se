import type { Route } from './route.types';

/**
 * Map of route definitions.
 */
const routes: Record<string, Route> = {};

/**
 * Stores a route definition in the routes map.
 */
export function defineRoute(route: Route): Route {
	routes[route.id] = route;
	if (!route.isActive) {
		route.isActive = (path, params) => path.startsWith(route.getPath(params));
	}
	return route;
}

/**
 * Returns a route definition by the route id
 */
export function getRoute(id: string) {
	return routes[id];
}

/**
 * Returns an array of all route definitions.
 */
export function getRoutes() {
	return Object.values(routes);
}

/**
 * Returns an array of child route definitions.
 */
export function getChildRoutes(id: string) {
	return getRoutes().filter((route) => route.parent?.id === id);
}

/**
 * Returns an array of active route definitions for a given path.
 */
export function getActiveRoutes(path: string, params: Record<string, string>): Route[] {
	return getRoutes()
		.filter((route) => route.isActive && route.isActive(path, params))
		.sort((a, b) => b.getPath(params).length - a.getPath(params).length);
}

/**
 * Returns routes that should be shown in the sidebar.
 */
export function getMenuItems(): Route[] {
	return getRoutes()
		.filter((route) => route.sidebar)
		.sort((a, b) => a.sidebar!.weight - b.sidebar!.weight);
}

/**
 * Returns the "homepage" route for an organization
 */
export function getHomeUrl(params: Record<string, string>) {
	const { organizationId } = params;
	if (organizationId) {
		return `/org/${organizationId}`;
	}
	return '/';
}

/**
 * Returns link to the parent route on the breadcrumb
 */
export function getParentUrl(route: Route, params: Record<string, string>) {
	if (route.parent) {
		return route.parent.getPath(params);
	}
	return getHomeUrl(params);
}

/**
 * Returns breadcrumbs for a given route definition
 */
export function getBreadcrumbs(route: Route): Route[] {
	const breadcrumbs: Route[] = [];
	let currentRoute: Route | null = route;
	while (currentRoute) {
		if (!currentRoute.hidden) {
			breadcrumbs.unshift(currentRoute);
		}
		currentRoute = currentRoute.parent ? routes[currentRoute.parent.id] : null;
	}
	return breadcrumbs;
}
