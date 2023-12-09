<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { page } from '$app/stores';
	import type { Project } from '$lib/types';
	import type { ActionResult } from '@sveltejs/kit';
	import {
		Breadcrumb,
		BreadcrumbItem,
		Button,
		Card,
		Dropdown,
		DropdownItem,
		Heading,
		Input,
		Label,
		Modal,
		Progressbar
	} from 'flowbite-svelte';
	import { DotsHorizontalOutline, ExclamationCircleOutline } from 'flowbite-svelte-icons';

	$: projects = $page.data.projects || [];
	$: projectListItems = projects.map((project: Project) => ({
		id: project.id,
		name: project.name,
		link: `/org/${project.organizationId}/project/${project.id}`
	}));

	let isCreateModalOpen = false;
	let isDeleteModalOpen = false;

	let projectToDelete: Project;

	let name = '';

	function selectProjectToDelete(project: Project) {
		projectToDelete = project;
		isDeleteModalOpen = true;
	}

	async function createProject() {
		const data = new FormData();
		data.set('name', name);
		const action = `/org/${$page.params.organizationId}/project?/createProject`;
		const response = await fetch(action, {
			method: 'POST',
			body: data
		});
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
	}

	async function deleteProject(projectId: string) {
		const data = new FormData();
		const action = `/org/${$page.params.organizationId}/project/${projectId}?/delete`;
		const response = await fetch(action, {
			method: 'POST',
			body: data
		});
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
	}
</script>

<Breadcrumb class="mb-6">
	<BreadcrumbItem href={`/org/${$page.params.organizationId}`} home>Home</BreadcrumbItem>
	<BreadcrumbItem>Projects</BreadcrumbItem>
</Breadcrumb>

<Heading customSize="mb-2 text-xl font-semibold">Projects</Heading>

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

<Modal title="Create Project" bind:open={isCreateModalOpen} autoclose>
	<Label for="name" class="block mb-2">Name:</Label>
	<Input id="name" required bind:value={name} />
	<svelte:fragment slot="footer">
		<Button size="sm" color="primary" on:click={createProject}>Create</Button>
	</svelte:fragment>
</Modal>

<Modal bind:open={isDeleteModalOpen} size="sm" autoclose>
	<div class="text-center">
		<ExclamationCircleOutline class="mx-auto mb-4 text-gray-500 w-8 h-8" />
		<h3 class="mb-5 text-lg font-normal text-gray-800">
			Are you sure you want to delete "{projectToDelete.name}"?
		</h3>
		<Button color="red" class="me-2" on:click={() => deleteProject(projectToDelete.id)}>
			Yes, I'm sure
		</Button>
		<Button color="alternative">No, cancel</Button>
	</div>
</Modal>

<style>
	.project-item {
		width: 300px;
		margin-right: 20px;
		margin-bottom: 20px;
	}
</style>
