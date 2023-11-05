<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Input, Label } from 'flowbite-svelte';
	import { CheckCircleOutline } from 'flowbite-svelte-icons';
	import { selectedOrg } from '$lib/stores/selected-org';

	let name = $selectedOrg?.name;
	let prod = $selectedOrg?.prod;

	async function submitForm(event: { currentTarget: EventTarget & HTMLFormElement }) {
		const data = new FormData();
		data.set('id', $selectedOrg.id);
		data.set('name', name);
		data.set('prod', prod ? 'true' : 'false');
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

<form method="post" action="?/updateOrganization" on:submit|preventDefault={submitForm}>
	<div class="mb-6">
		<Label for="name" class="block mb-2">Organization Name:</Label>
		<Input id="name" required bind:value={name} />
	</div>

	<div>
		<Button type="submit" size="xs" color="dark">
			Update organization &nbsp;
			<CheckCircleOutline class="w-3 h-3 mr-1" />
		</Button>
	</div>
</form>
