<script lang="ts">
	import { ChevronRightIcon } from 'lucide-svelte';
	import { type Project, ProjectStatus } from '$lib/types';

	const ORDERED_STATUSES = [
		{ value: ProjectStatus.Collecting, label: 'Collecting' },
		{ value: ProjectStatus.Backlog, label: 'Backlog' },
		{ value: ProjectStatus.Completed, label: 'Completed' }
	];

	type GroupedProjects = {
		[key: string]: Project[];
	};

	export let projects: Project[] = [];

	$: groups = projects.reduce((acc: GroupedProjects, project) => {
		if (!acc[project.status]) {
			acc[project.status] = [];
		}
		acc[project.status].push(project);
		return acc;
	}, {});
</script>

{#each ORDERED_STATUSES as status}
	{#if groups[status.value]?.length}
		<div class="divider">{status.label}</div>
		<table class="table">
			<thead>
				<tr>
					<th>Name</th>
					<th>Type</th>
					<th>Created At</th>
					<th></th>
				</tr>
			</thead>
			<tbody>
				{#each groups[status.value] as project}
					<tr>
						<td>{project.name}</td>
						<td>{project.feedbackType}</td>
						<td class="hidden md:table-cell">{project.createdAt.toLocaleDateString()}</td>
						<td class="text-right">
							<a href="./projects/{project.id}" role="button" class="btn btn-sm btn-ghost">
								<ChevronRightIcon class="w-4 h-4" />
							</a>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	{/if}
{/each}
