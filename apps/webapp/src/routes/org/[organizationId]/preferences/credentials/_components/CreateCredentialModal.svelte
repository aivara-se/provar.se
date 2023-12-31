<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Input, Label, Modal } from 'flowbite-svelte';

	export let isOpen: boolean = false;

	interface Fields {
		name: string;
	}

	let fields: Fields = {
		name: ''
	};

	function closeModal() {
		isOpen = false;
	}

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
		const action = '?/createCredential';
		const response = await fetch(action, { method: 'post', body: data });
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
		resetFields();
		closeModal();
	}
</script>

<Modal title="Create API Key" bind:open={isOpen} autoclose>
	<form method="post" on:submit|self|stopPropagation|preventDefault={handler}>
		<Label for="name" class="block mb-2">Name:</Label>
		<Input id="name" required bind:value={fields.name} />
	</form>
	<svelte:fragment slot="footer">
		<Button type="submit" size="sm" color="primary" on:click={handler}>Create</Button>
	</svelte:fragment>
</Modal>
