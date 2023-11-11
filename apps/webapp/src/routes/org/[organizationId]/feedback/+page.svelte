<script lang="ts">
	import { Button } from 'flowbite-svelte';
	import { page } from '$app/stores';
	import type { Project } from '$lib/types';

	$: createLink = `/org/${$page.params.organizationId}/feedback/project/create`;

	$: projectListItems = $page.data.projects.map((project: Project) => ({
		name: project.name,
		link: `/org/${project.organizationId}/feedback/project/${project.id}`
	}));
</script>

<Button href={createLink}>Create Project</Button>

<ul>
	{#each projectListItems as project}
		<li>
			<a href={project.link}>{project.name}</a>
		</li>
	{/each}
</ul>
