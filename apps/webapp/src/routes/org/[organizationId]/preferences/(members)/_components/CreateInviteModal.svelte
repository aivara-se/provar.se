<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Input, Label, Modal } from 'flowbite-svelte';

	export let isOpen: boolean = false;

	interface Fields {
		name: string;
		email: string;
	}

	let fields: Fields = {
		name: '',
		email: ''
	};

	function closeModal() {
		isOpen = false;
	}

	function resetFields() {
		fields = {
			name: '',
			email: ''
		};
	}

	async function handler() {
		const data = new FormData();
		for (const [key, value] of Object.entries(fields)) {
			data.set(key, value);
		}
		const action = `?/inviteMember`;
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

<Modal title="Invite Member" bind:open={isOpen} autoclose>
	<form method="post" on:submit|self|stopPropagation|preventDefault={handler}>
		<section class="flex flex-col gap-6">
			<div>
				<Label for="name" class="block mb-1">Name:</Label>
				<Input id="name" required bind:value={fields.name} />
			</div>
			<div>
				<Label for="email" class="block mb-1">Email:</Label>
				<Input id="email" required bind:value={fields.email} />
			</div>
		</section>
	</form>
	<svelte:fragment slot="footer">
		<Button type="submit" size="sm" color="primary" on:click={handler}>Invite</Button>
	</svelte:fragment>
</Modal>
