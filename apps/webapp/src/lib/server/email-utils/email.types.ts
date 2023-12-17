/**
 * Email templates and details needed to render and send an email
 */
export interface EmailDetails<T> {
	toEmail: string;
	options: T;
	template: {
		subject: (options: T) => string;
		text: (options: T) => string;
		html: (options: T) => string;
	};
}
