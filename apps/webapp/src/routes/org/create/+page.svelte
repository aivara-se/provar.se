<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { CenteredLayout } from '$lib/ui';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Input, Label } from 'flowbite-svelte';
	import { UsersGroupOutline } from 'flowbite-svelte-icons';

	let name = '';

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

<CenteredLayout>
	<form method="post" on:submit|preventDefault={submitForm}>
		<div class="grid gap-6 mb-6">
			<div>
				<Label color="gray" for="name" class="mb-2">Organization name</Label>
				<Input type="text" id="name" required bind:value={name} />
			</div>
			<div>
				<Button outline type="submit" size="sm">
					Create organization &nbsp;
					<UsersGroupOutline class="w-3 h-3 mr-1" />
				</Button>
			</div>
		</div>
	</form>
</CenteredLayout>
