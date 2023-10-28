<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { CenteredLayout } from '$lib/ui';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Input, Label } from 'flowbite-svelte';

	let name = '';

	async function submitForm(event: { currentTarget: EventTarget & HTMLFormElement }) {
		const data = new FormData();
		data.set('name', name);
		const response = await fetch(event.currentTarget.action, {
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

<CenteredLayout>
	<form method="POST" on:submit|preventDefault={submitForm}>
		<div class="grid gap-6 mb-6">
			<div>
				<Label color="gray" for="name" class="mb-2">Project name</Label>
				<Input type="text" id="name" required bind:value={name} />
			</div>
			<div>
				<Button outline type="submit" size="sm">Create Project</Button>
			</div>
		</div>
	</form>
</CenteredLayout>
