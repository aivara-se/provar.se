import { ProjectRepository } from '$lib/server';
import type { Actions } from './$types';
import type { Action } from '@sveltejs/kit';
import type { RequestEvent } from './$types';
import type { RouteParams } from './$types';

export const actions: Actions = {
	updateProject: withProject(async (project, { request }) => {
		const data = await request.formData();
		const name = String(data.get('name'));
		await ProjectRepository.update(project.id, { name });
	})
};

export function withProject(
	arg0: (project: any, { request }: { request: any }) => Promise<void>
): Action<
	RouteParams,
	void | Record<string, unknown>,
	'/org/[organizationId]/feedback/project/[projectId]'
> {
	return async function (this: RequestEvent, { params }) {
		try {
			const projectId = params.projectId;
			const response = await fetch(`/api/projects/${projectId}`);
			const project = await response.json();

			await arg0(project, { request: this.request });
		} catch (error) {
			console.error('Error:', error);
		}
	};
}
