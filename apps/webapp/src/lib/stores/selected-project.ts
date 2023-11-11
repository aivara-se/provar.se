import { page } from '$app/stores';
import { derived } from 'svelte/store';
import type { Project } from '../types';

/**
 * This is a derived store that returns the selected projects from the
 * current page's data. Returns null if outside a project page.
 */
export const selectedProject = derived(page, ($page) => {
	const projectId = $page.params.organizationId;
	const project: Project[] = $page.data.project;
	if (!projectId || !project) {
		return null;
	}
	return project.find((project) => project.id === projectId);
});
