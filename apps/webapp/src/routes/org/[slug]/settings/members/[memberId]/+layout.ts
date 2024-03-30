import { error } from '@sveltejs/kit';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async (event) => {
	const { members } = await event.parent();
	const member = members.find((user) => user.id === event.params.memberId);
	if (!member) {
		error(404);
	}
	return { member };
};
