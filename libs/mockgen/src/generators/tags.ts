import { POSSIBLE_TAGS } from './tags.data.js';

export function generateRandomTags() {
	const tags: Record<string, string> = {};
	for (const [key, fn] of Object.entries(POSSIBLE_TAGS)) {
		tags[key] = fn();
	}
	return tags;
}
