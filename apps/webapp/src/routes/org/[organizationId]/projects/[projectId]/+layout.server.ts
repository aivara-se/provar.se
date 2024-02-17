import { error } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async (event) => {
	const { projects } = await event.parent();
	const project = projects.find((cred) => cred.id === event.params.projectId);
	if (!project) {
		error(404);
	}
	return { project };
};
