<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Input, Label } from 'flowbite-svelte';
	import { CheckCircleOutline } from 'flowbite-svelte-icons';
	import { selectedProject } from '$lib/stores/selected-project';

	let name = $selectedProject?.name || '';

	async function submitForm(event: { currentTarget: EventTarget & HTMLFormElement }) {
		const data = new FormData();
		data.set('name', name);
		const response = await fetch(event.currentTarget.action, {
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

<form method="post" action="?/updateProject" on:submit|preventDefault={submitForm}>
	<div class="mb-6">
		<Label for="name" class="block mb-2">Project Name:</Label>
		<Input id="name" required bind:value={name} />
	</div>

	<div>
		<Button type="submit" size="xs" color="dark">
			Update Project &nbsp;
			<CheckCircleOutline class="w-3 h-3 mr-1" />
		</Button>
	</div>
</form>

<!--
	<script>
	import { Alert } from 'flowbite-svelte';
	import { InfoCircleSolid } from 'flowbite-svelte-icons';
</script>

<Alert color="blue" rounded={false} class="border-t-4">
	<InfoCircleSolid slot="icon" class="w-4 h-4" />
	<span class="font-medium">Project details</span><br />
	This page is under construction.
</Alert>

-->
