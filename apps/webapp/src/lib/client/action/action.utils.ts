import { getActiveRoutes } from '$lib/client/routes';
import type { Action, ActionGroup } from './action.types';

/**
 * Returns an array of actions for route
 */
export function getActions(path: string, params: Record<string, string>): ActionGroup[] {
	const routes = getActiveRoutes(path, params).filter((route) => route.getActions);
	const actions = routes.flatMap((route) => route.getActions!(path, params));
	return groupActions(actions);
}

/**
 * Groups actions by their group label.
 */
function groupActions(actions: Action[]): ActionGroup[] {
	return actions.reduce((groups, action) => {
		const group = groups.find((group) => group.name === action.group);
		if (group) {
			group.actions.push(action);
		} else {
			groups.push({ name: action.group, actions: [action] });
		}
		return groups;
	}, [] as ActionGroup[]);
}

/**
 * Performs an action based on the action type
 */
export async function performAction(action: Action) {
	if (action.type === 'frontend') {
		await action.exec();
	}
}
