import type { Provider } from '@auth/core/providers';
import * as VerifyLoginTemplate from './templates/verifylogin.tpl';
import * as EmailService from './email.service';

export function EmailProvider(): Provider {
	return {
		id: 'email',
		type: 'email',
		name: 'EmailProvider',
		from: '',
		maxAge: 24 * 60 * 60,
		options: {},
		async sendVerificationRequest(params) {
			await EmailService.send({
				toEmail: params.identifier,
				options: { link: params.url },
				template: VerifyLoginTemplate
			});
		}
	};
}
