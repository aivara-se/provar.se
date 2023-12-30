import { USER_AGENTS } from './meta.data.js';

function generateIpAddress(): string {
	const r = () => Math.floor(Math.random() * 256);
	return `${r()}.${r()}.${r()}.${r()}`;
}

function generateUserAgent(): string {
	return USER_AGENTS[Math.floor(Math.random() * USER_AGENTS.length)];
}

export function generateMetadata(): Record<string, unknown> {
	const meta: Record<string, unknown> = {};
	meta['request.ip'] = Math.random() > 0.2 ? generateIpAddress() : '';
	meta['request.header.user-agent'] = Math.random() > 0.2 ? generateUserAgent() : '';
	return meta;
}
