<script lang="ts">
	import { api } from '$lib/client/api';
	import { ConfirmModal } from '$lib/client/forms';
	import type { Invitation } from '$lib/client/types';

	export let invitation: Invitation;

	let isOpen = false;

	async function resendInvitation() {
		// await api().Invitation.resend(invitation.organizationId, invitation.id);
	}
</script>

<section class="mt-8 rounded-lg p-4 bg-gray-100 text-gray-900 dark:bg-gray-950 dark:text-gray-200">
	<h3 class="flex items-center justify-between text-lg font-medium">Resend invitation</h3>
	<p class="text-sm mt-2">
		Resend the invitation if the recipient didn't receive it or if the link is no longer available.
		The user will receive a new email with the invitation link.
	</p>
	<button class="btn btn-sm btn-neutral mt-4" on:click={() => (isOpen = true)}>
		Resend invitation
	</button>
	<ConfirmModal bind:isOpen action={resendInvitation} submitText="Resend invitation">
		<p class="py-4 text-center">Are you sure you want to resend the invitation?</p>
	</ConfirmModal>
</section>
