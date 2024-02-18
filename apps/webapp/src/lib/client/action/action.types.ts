interface BaseAction {
	id: string;
	type: 'navigate' | 'backend' | 'frontend';
	name: string;
	group?: string;
}

export interface NavigateAction extends BaseAction {
	type: 'navigate';
	path: string;
}

export interface BackendAction extends BaseAction {
	type: 'backend';
	path: string;
	formAction: string;
}

export interface FrontentAction extends BaseAction {
	type: 'frontend';
	exec: () => void | Promise<void>;
}

export type Action = NavigateAction | BackendAction | FrontentAction;

export interface ActionGroup {
	name?: string;
	actions: Action[];
}
