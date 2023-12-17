import { ProjectRepository } from '$lib/server/project';
import type { LayoutServerLoad } from './$types';

/**
 * Loads data required for all tabs on organization preferences page.
 */
export const load: LayoutServerLoad = async (event) => {
	const { organization } = await event.parent();
	const projects = await ProjectRepository.findByOrganization(organization.id);
	return { projects: projects };
};
