<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { CopyableCode } from '$lib/client/ui';
	import type { User } from '$lib/types';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Heading, Modal } from 'flowbite-svelte';

	export let isOpen: boolean = false;
	export let member: User | null = null;

	function closeModal() {
		isOpen = false;
	}

	async function handler() {
		if (!member) {
			return;
		}
		const data = new FormData();
		data.set('id', member.id);
		const action = `?/revokeMembership`;
		const response = await fetch(action, { method: 'post', body: data });
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
		closeModal();
	}
</script>

<Modal title={member?.name || ''} bind:open={isOpen} autoclose>
	<section>
		<Heading customSize="font-semibold mb-2">Email:</Heading>
		<CopyableCode text={member?.email} />
	</section>
	<svelte:fragment slot="footer">
		<Button size="sm" color="red" on:click={handler}>Revoke</Button>
	</svelte:fragment>
</Modal>
