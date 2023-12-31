<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import type { Organization } from '$lib/types';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Modal, P } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';

	export let isOpen: boolean = false;
	export let organization: Organization;

	function closeModal() {
		isOpen = false;
	}

	async function handler() {
		const action = '?/deleteOrganization';
		const response = await fetch(action, { method: 'post', body: new FormData() });
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
		closeModal();
	}
</script>

<Modal size="sm" bind:open={isOpen} autoclose>
	<div class="text-center">
		<ExclamationCircleOutline class="mx-auto mb-4 text-gray-500 w-8 h-8" />
		<h3 class="mb-5 text-lg font-normal text-gray-800">
			Are you sure you want to delete "{organization.name}"?
		</h3>
		<Button color="red" class="me-2" on:click={handler}>Yes, I'm sure</Button>
		<Button color="alternative" on:click={closeModal}>No, cancel</Button>
	</div>
</Modal>
