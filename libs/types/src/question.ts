export interface Question {
	id: string;
	organization_id: string;
	type: string;
	text: string;
	createdAt: string;
	updatedAt: string;
}

export interface TextQuestion extends Question {
	type: 'text';
}

export interface NPSQuestion extends Question {
	type: 'nps';
}

export interface CSATQuestion extends Question {
	type: 'csat';
}
