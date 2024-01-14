<script lang="ts">
	import { page } from '$app/stores';
	import type { Project } from '$lib/types';
	import {
		Breadcrumb,
		BreadcrumbItem,
		Button,
		Card,
		Dropdown,
		DropdownItem,
		Progressbar
	} from 'flowbite-svelte';
	import { DotsHorizontalOutline } from 'flowbite-svelte-icons';
	import CreateProjectModel from './_components/CreateProjectModel.svelte';
	import DeleteProjectModal from './_components/DeleteProjectModal.svelte';

	$: projects = $page.data.projects || [];
	$: projectListItems = projects.map((project: Project) => ({
		id: project.id,
		name: project.name,
		link: `/org/${project.organizationId}/project/${project.id}`
	}));

	let isCreateModalOpen = false;
	let isDeleteModalOpen = false;

	let projectToDelete: Project;

	function selectProjectToDelete(project: Project) {
		projectToDelete = project;
		isDeleteModalOpen = true;
	}
</script>

<section class="flex items-start h-10 justify-between mb-1">
	<section>
		<Breadcrumb class="mb-6">
			<BreadcrumbItem href={`/org/${$page.params.organizationId}`} home>Home</BreadcrumbItem>
			<BreadcrumbItem>Projects</BreadcrumbItem>
		</Breadcrumb>
	</section>
	<section class="flex gap-2 items-start"></section>
</section>

<section>
	{#each projectListItems as item}
		<div class="inline-block mb-6 project-item">
			<Card shadow={false}>
				<div class="mb-2 flex items-center justify-between">
					<a href={item.link}>
						<h5 class="mb-2 font-semibold text-gray-900">{item.name}</h5>
					</a>
					<div class="p-2 rounded hover:bg-gray-300">
						<DotsHorizontalOutline class="w-3 h-3 text-gray-700 outline-none" />
						<Dropdown class="w-36">
							<DropdownItem>Modify</DropdownItem>
							<DropdownItem on:click={() => selectProjectToDelete(item)}>Delete</DropdownItem>
						</Dropdown>
					</div>
				</div>
				<Progressbar progress="30" />
			</Card>
		</div>
	{/each}
</section>

<section>
	<Button color="light" size="xs" on:click={() => (isCreateModalOpen = true)}>
		+ Create project
	</Button>
</section>

<CreateProjectModel bind:isOpen={isCreateModalOpen} />
<DeleteProjectModal bind:isOpen={isDeleteModalOpen} project={projectToDelete} />

<style>
	.project-item {
		width: 300px;
		margin-right: 20px;
		margin-bottom: 20px;
	}
</style>
