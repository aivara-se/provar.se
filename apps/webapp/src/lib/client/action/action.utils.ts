import { getActiveRoutes } from '$lib/client/routes';
import type { Action, ActionGroup } from './action.types';

/**
 * Returns an array of actions for route
 */
export function getActions(path: string, params: Record<string, string>): ActionGroup[] {
	const routes = getActiveRoutes(path, params).filter((route) => route.getActions);
	const actions = routes.flatMap((route) => route.getActions!(path, params));
	const groups = groupActions(actions).map((group) => sortGroupActions(group));
	return groups;
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
 * Sorts actions in a group by their name in-place.
 */
function sortGroupActions(group: ActionGroup): ActionGroup {
	group.actions.sort((a, b) => (a.name > b.name ? 1 : -1));
	return group;
}

/**
 * Performs an action based on the action type
 */
export async function performAction(action: Action) {
	if (action.type === 'frontend') {
		await action.exec();
	}
}
