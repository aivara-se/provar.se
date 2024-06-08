@discriminator("type")
model Answer {
	id: string;
	organization_id: string;
	question_id: string;
	response_id: string;
	type: string;
	value: string;
	createdAt: string;
	updatedAt: string;
}

model TextAnswer extends Answer {
	type: "text";
	value: string;
}

model ElevenPointsAnswer extends Answer {
	type: "11pts";
	value: "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" | "10";
}

model FivePointsAnswer extends Answer {
	type: "5pts";
	value: "0" | "1" | "2" | "3" | "4" | "5";
}

model ThreePointsAnswer extends Answer {
	type: "3pts";
	value: "0" | "1" | "2";
}

model YesNoAnswer extends Answer {
	type: "yesno";
	value: "0" | "1";
}
