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
		Heading,
		Input,
		Label,
		Modal,
		Progressbar
	} from 'flowbite-svelte';

	$: projects = $page.data.projects || [];
	$: projectListItems = projects.map((project: Project) => ({
		name: project.name,
		link: `/org/${project.organizationId}/project/${project.id}`
	}));

	let isCreateModalOpen = false;

	let name = '';

	async function createProject() {
		const data = new FormData();
		data.set('name', name);
		const action = `/org/${$page.params.organizationId}/project`;
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

<style>
	.project-item {
		width: 300px;
		margin-right: 20px;
		margin-bottom: 20px;
	}
</style>
