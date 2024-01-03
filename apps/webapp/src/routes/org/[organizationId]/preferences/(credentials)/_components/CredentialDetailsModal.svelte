<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { CopyableCode } from '$lib/client/ui';
	import type { Credential } from '$lib/types';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Heading, Modal } from 'flowbite-svelte';

	export let isOpen: boolean = false;
	export let credential: Credential;

	function closeModal() {
		isOpen = false;
	}

	async function handler() {
		const data = new FormData();
		data.set('id', credential.id);
		const action = `?/revokeCredential`;
		const response = await fetch(action, { method: 'post', body: data });
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
		closeModal();
	}
</script>

<Modal title={credential.name} bind:open={isOpen} autoclose>
	<main>
		<Heading customSize="font-semibold mb-2">API Key:</Heading>
		<CopyableCode text={credential.key} />
	</main>
	<svelte:fragment slot="footer">
		<Button size="sm" color="red" on:click={handler}>Revoke</Button>
	</svelte:fragment>
</Modal>
