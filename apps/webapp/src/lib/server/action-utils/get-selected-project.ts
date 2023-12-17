import { ProjectRepository } from '$lib/server/project';
import { error, type RequestEvent } from '@sveltejs/kit';
import { getSelectedOrganization } from './get-selected-organization';

/**
 * A helper function to get the selected project using path params.
 * Expects "projectId" to be present in the path params.
 */
export async function getSelectedProject(event: RequestEvent) {
	const organization = await getSelectedOrganization(event);
	if (!event.params.projectId) {
		throw error(404);
	}
	const project = await ProjectRepository.findById(organization.id, event.params.projectId);
	if (!project) {
		throw error(404);
	}
	return { organization, project };
}
