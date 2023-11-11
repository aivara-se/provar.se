<script lang="ts">
	import { Button } from 'flowbite-svelte';
	import { page } from '$app/stores';
	import type { Project } from '$lib/types';

	$: createLink = `/org/${$page.params.organizationId}/feedback/project/create`;

	function transformToListItem(project: Project) {
		return {
			name: project.name,
			link: `/org/${project.organizationId}/feedback/project/${project.id}`
		};
	}

	$: projectListItems = $page.data.projects.map(transformToListItem);
</script>

<Button href={createLink}>Create Project</Button>

<ul>
	{#each projectListItems as project}
		<li>
			<a href={project.link}>{project.name}</a>
		</li>
	{/each}
</ul>
