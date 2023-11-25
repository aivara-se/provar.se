<script lang="ts">
	import { page } from '$app/stores';
	import type { Project } from '$lib/types';
	import { Breadcrumb, BreadcrumbItem, Button, Card, Heading, Progressbar } from 'flowbite-svelte';

	$: projects = $page.data.projects || [];
	$: projectListItems = projects.map((project: Project) => ({
		name: project.name,
		link: `/org/${project.organizationId}/feedback/project/${project.id}`
	}));

	$: createLink = `/org/${$page.params.organizationId}/feedback/project/create`;
</script>

<Breadcrumb class="mb-6">
	<BreadcrumbItem href={`/org/${$page.params.organizationId}`} home>Home</BreadcrumbItem>
	<BreadcrumbItem href={`/org/${$page.params.organizationId}/feedback`}>Feedbacks</BreadcrumbItem>
	<BreadcrumbItem>Projects</BreadcrumbItem>
</Breadcrumb>

<Heading customSize="mb-2 text-xl font-semibold">Projects</Heading>

<section>
	{#each projectListItems as item}
		<div class="inline-block mb-6 project-item">
			<Card href={item.link} shadow={false}>
				<h5 class="mb-2 font-semibold text-gray-900">
					{item.name}
				</h5>
				<Progressbar progress="50" />
			</Card>
		</div>
	{/each}
</section>

<section>
	<Button href={createLink} color="light" size="sm" outline>Create new project</Button>
</section>

<style>
	.project-item {
		width: 300px;
		margin-right: 20px;
		margin-bottom: 20px;
	}
</style>
