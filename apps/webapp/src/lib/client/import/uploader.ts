import { api } from '../api';

/**
 * Uplaoder uploads given feedback data to the server
 */
export async function upload(orgId: string, data: Record<string, string>) {
	await api().Feedback.create(orgId, {
		feedbackType: getFeedbackType(data),
		feedbackTime: getFeedbackTime(data),
		feedbackData: getFeedbackData(data),
		feedbackTags: getFeedbackTags(data),
		feedbackUser: getFeedbackUser(data),
		feedbackMeta: getFeedbackMeta(data)
	});
}

function getFeedbackType(data: Record<string, string>): 'text' | 'cnps' | 'csat' {
	return data['type'] as 'text' | 'cnps' | 'csat';
}

function getFeedbackTime(data: Record<string, string>): string {
	return data['time'] as string;
}

function getFeedbackData(data: Record<string, string>) {
	return withPrefix('data', data);
}

function getFeedbackTags(data: Record<string, string>) {
	return withPrefix('tags', data);
}

function getFeedbackUser(data: Record<string, string>) {
	return withPrefix('user', data);
}

function getFeedbackMeta(data: Record<string, string>) {
	return withPrefix('meta', data);
}

function withPrefix(prefix: string, data: Record<string, string>) {
	const feedbackData: Record<string, string> = {};
	for (const key in data) {
		if (key.startsWith(`${prefix}.`)) {
			feedbackData[key.substring(prefix.length + 1)] = data[key];
		}
	}
	return feedbackData;
}
