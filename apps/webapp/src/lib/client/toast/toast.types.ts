/**
 * Describes a notification which expires after a duration
 */
export interface Toast {
	id: string;
	type: 'success' | 'info' | 'warning' | 'error';
	content: string;
	duration?: number;
}
