<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Input, Label } from 'flowbite-svelte';
	import { UsersGroupOutline } from 'flowbite-svelte-icons';

	interface Fields {
		name: string;
	}

	let fields: Fields = {
		name: ''
	};

	function resetFields() {
		fields = {
			name: ''
		};
	}

	async function handler() {
		const data = new FormData();
		for (const [key, value] of Object.entries(fields)) {
			data.set(key, value);
		}
		const response = await fetch('', { method: 'post', body: data });
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
		resetFields();
	}
</script>

<section>
	<form on:submit|self|stopPropagation|preventDefault={handler}>
		<div class="grid gap-6 mb-6">
			<div>
				<Label color="gray" for="name" class="mb-2">Organization name</Label>
				<Input type="text" id="name" required bind:value={fields.name} />
			</div>
			<div>
				<Button outline type="submit" size="sm">
					Create organization &nbsp;
					<UsersGroupOutline class="w-3 h-3 mr-1" />
				</Button>
			</div>
		</div>
	</form>
</section>
