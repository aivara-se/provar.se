import { error } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

/**
 * Loads data required for all tabs on organization preferences page.
 */
export const load: LayoutServerLoad = async (event) => {
	const { projects } = await event.parent();
	const project = projects.find((project) => project.id === event.params.projectId);
	if (!project) {
		error(403);
	}
	return { project };
};
