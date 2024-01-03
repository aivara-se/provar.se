<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { CopyableCode } from '$lib/client/ui';
	import type { Invitation } from '$lib/types';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Heading, Modal } from 'flowbite-svelte';

	export let isOpen: boolean = false;
	export let invitation: Invitation | null = null;

	function closeModal() {
		isOpen = false;
	}

	async function revokeInvitation() {
		if (!invitation) {
			return;
		}
		const data = new FormData();
		data.set('id', invitation.id);
		const action = `?/revokeInvitation`;
		const response = await fetch(action, { method: 'post', body: data });
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
		closeModal();
	}

	async function resendInvitation() {
		if (!invitation) {
			return;
		}
		const data = new FormData();
		data.set('id', invitation.id);
		const action = `?/resendInvitation`;
		const response = await fetch(action, { method: 'post', body: data });
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
		closeModal();
	}
</script>

<Modal title={invitation?.name || ' '} bind:open={isOpen} autoclose>
	<main>
		<section>
			<Heading customSize="font-semibold mb-2">Email:</Heading>
			<CopyableCode text={invitation?.email} />
		</section>
		<section class="mt-3">
			<Heading customSize="font-semibold mb-2">Link:</Heading>
			<CopyableCode text={invitation?.link} />
		</section>
	</main>
	<svelte:fragment slot="footer">
		<Button size="sm" color="red" on:click={revokeInvitation}>Revoke</Button>
		<Button size="sm" color="dark" on:click={resendInvitation}>Resend</Button>
	</svelte:fragment>
</Modal>
