<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import type { Credential } from '$lib/types';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Modal, P } from 'flowbite-svelte';

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

<Modal title={credential.name} size="sm" bind:open={isOpen} autoclose>
	<P class="mb-2">
		<span class="font-semibold">API Key:</span>
		<code class="block break-words text-xs mt-2 bg-gray-100 p-2 rounded">
			{credential.key}
		</code>
	</P>
	<svelte:fragment slot="footer">
		<Button size="sm" color="red" on:click={handler}>Revoke</Button>
	</svelte:fragment>
</Modal>
