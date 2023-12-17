import nodemailer from 'nodemailer';
import { env } from '$env/dynamic/private';
import type { EmailDetails } from './email.types';

/**
 * Create a nodemailer transport which can be used to send emails
 */
const transport = nodemailer.createTransport(env.AUTH_EMAIL_SERVER);

/**
 * Send an email using given templates and details
 */
export async function send<T>(email: EmailDetails<T>) {
	await transport.sendMail({
		from: env.AUTH_EMAIL_FROM,
		to: email.toEmail,
		subject: email.template.subject(email.options),
		text: email.template.text(email.options),
		html: email.template.html(email.options)
	});
}
