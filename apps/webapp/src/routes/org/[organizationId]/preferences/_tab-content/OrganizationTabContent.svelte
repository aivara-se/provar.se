<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { selectedOrg } from '$lib/stores/selected-org';
	import type { ActionResult } from '@sveltejs/kit';
	import { Alert, Button, Heading, Input, Label, Modal, P, Textarea } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';

	let name = $selectedOrg?.name || '';
	let description = $selectedOrg?.description || '';

	let isDeleteModalOpen = false;

	async function deleteOrganization() {
		const data = new FormData();
		const response = await fetch('?/deleteOrganization', {
			method: 'post',
			body: data
		});
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
	}

	async function updateOrganization() {
		const data = new FormData();
		data.set('name', name);
		data.set('description', description);
		const response = await fetch('?/updateOrganization', {
			method: 'post',
			body: data
		});
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
	}
</script>

<Heading customSize="text-xl font-semibold">Organization details</Heading>

<P class="mt-6">
	You can update your organization's information here to ensure accuracy and relevance which will be
	used by AI features.
</P>

<div class="mt-6">
	<Label for="name" class="block mb-2">Name:</Label>
	<Input id="name" required bind:value={name} />
</div>

<div class="mt-6">
	<Label for="description" class="block mb-2">Description:</Label>
	<Textarea id="description" required bind:value={description} />
</div>

<Alert color="red" class="mt-6">
	<div class="flex items-center gap-3">
		<ExclamationCircleOutline slot="icon" class="w-4 h-4" />
		<span class="font-medium">Delete the organization</span>
	</div>
	<p class="mt-2 mb-4 text-sm">
		Deleting this organization will permanently erase all collected feedback data, including
		associated records and information. Proceeding with this action cannot be undone. Please ensure
		that you have backed up any critical data or exported necessary information before deleting the
		organization.
	</p>
	<div class="flex gap-2">
		<Button size="xs" color="red" outline on:click={() => (isDeleteModalOpen = true)}>
			Delete "{name}"
		</Button>
	</div>
</Alert>

<div class="flex justify-end mt-6">
	<Button size="sm" color="primary" on:click={updateOrganization}>Update</Button>
</div>

<Modal bind:open={isDeleteModalOpen} size="sm" autoclose>
	<div class="text-center">
		<ExclamationCircleOutline class="mx-auto mb-4 text-gray-500 w-8 h-8" />
		<h3 class="mb-5 text-lg font-normal text-gray-800">
			Are you sure you want to delete "{name}"?
		</h3>
		<Button color="red" class="me-2" on:click={deleteOrganization}>Yes, I'm sure</Button>
		<Button color="alternative">No, cancel</Button>
	</div>
</Modal>
