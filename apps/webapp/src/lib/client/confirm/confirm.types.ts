/**
 * Describes a notification which expires after a duration
 */
export interface Toast {
	type: 'success' | 'info' | 'warning' | 'error';
	content: string;
	duration?: number;
}

/**
 * Describes a confirm dialog with callbacks to be executed
 */
export interface Confirm {
	title: string;
	content: string;
	onSubmit?: () => void;
	onCancel?: () => void;
}
