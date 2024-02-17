import { error } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async (event) => {
	const { members } = await event.parent();
	const member = members.find((user) => user.id === event.params.memberId);
	if (!member) {
		error(404);
	}
	return { member };
};
