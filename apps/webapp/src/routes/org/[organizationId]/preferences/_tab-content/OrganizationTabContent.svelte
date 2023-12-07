<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { selectedOrg } from '$lib/stores/selected-org';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Input, Label, Textarea } from 'flowbite-svelte';

	let name = $selectedOrg?.name || '';
	let description = $selectedOrg?.description || '';

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

<div class="mb-6">
	<Label for="name" class="block mb-2">Name:</Label>
	<Input id="name" required bind:value={name} />
</div>

<div class="mb-6">
	<Label for="description" class="block mb-2">Description:</Label>
	<Textarea id="description" required bind:value={description} />
</div>

<div class="flex justify-between">
	<Button size="sm" color="red" outline on:click={deleteOrganization}>Delete</Button>
	<Button size="sm" color="primary" on:click={updateOrganization}>Update</Button>
</div>
