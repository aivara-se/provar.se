model Question {
	id: string;
	organization_id: string;
	type: string;
	text: string;
	createdAt: string;
	updatedAt: string;
}

model TextQuestion extends Question {
	type: "text";
}

model NPSQuestion extends Question {
	type: "nps";
}

model CSATQuestion extends Question {
	type: "csat";
}
