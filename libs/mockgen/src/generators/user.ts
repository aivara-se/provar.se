import { EMAIL_DOMAINS, FIRST_NAMES, LAST_NAMES } from './user.data.js';

export interface User {
	id: string;
	name: string;
	email: string;
	avatar: string;
}

function generateFirstName(): string {
	return FIRST_NAMES[Math.floor(Math.random() * FIRST_NAMES.length)];
}

function generateLastName(): string {
	return LAST_NAMES[Math.floor(Math.random() * LAST_NAMES.length)];
}

function generateName(): string {
	return `${generateFirstName()} ${generateLastName()}`;
}

function generateEmail(name: string): string {
	const user = name
		.toLowerCase()
		.replace(' ', '.')
		.replace(/[^a-zA-Z0-9]/g, '');
	const host = EMAIL_DOMAINS[Math.floor(Math.random() * EMAIL_DOMAINS.length)];
	return `${user}@${host}`;
}

export function generateUserdata(): User {
	const id = Math.random().toString(36).substr(2, 9);
	const name = generateName();
	const email = generateEmail(name);
	const avatar = `https://i.pravatar.cc/150?u=${id}`;
	return { id, name, email, avatar };
}
