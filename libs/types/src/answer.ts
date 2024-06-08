export interface Answer {
	id: string;
	organization_id: string;
	question_id: string;
	response_id: string;
	type: string;
	value: string;
	createdAt: string;
	updatedAt: string;
}

export interface TextAnswer extends Answer {
	type: 'text';
	value: string;
}

export interface ElevenPointsAnswer extends Answer {
	type: '11pts';
	value: '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9' | '10';
}

export interface FivePointsAnswer extends Answer {
	type: '5pts';
	value: '0' | '1' | '2' | '3' | '4' | '5';
}

export interface ThreePointsAnswer extends Answer {
	type: '3pts';
	value: '0' | '1' | '2';
}

export interface YesNoAnswer extends Answer {
	type: 'yesno';
	value: '0' | '1';
}
