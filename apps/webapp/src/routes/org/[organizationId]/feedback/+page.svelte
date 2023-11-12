<script lang="ts">
	import { Button } from 'flowbite-svelte';
	import { page } from '$app/stores';
	import type { Project } from '$lib/types';

	$: projects = $page.data.projects || [];
	$: projectListItems = projects.map((project: Project) => ({
		name: project.name,
		link: `/org/${project.organizationId}/feedback/project/${project.id}`
	}));

	$: createLink = `/org/${$page.params.organizationId}/feedback/project/create`;
</script>

<Button href={createLink}>Create Project</Button>

<ul>
	{#each projectListItems as project}
		<li>
			<a href={project.link}>{project.name}</a>
		</li>
	{/each}
</ul>
