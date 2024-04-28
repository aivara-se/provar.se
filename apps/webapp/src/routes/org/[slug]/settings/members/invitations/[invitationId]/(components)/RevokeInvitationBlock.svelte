<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/client/api';
	import { ConfirmModal } from '$lib/client/forms';
	import type { Invitation, Organization } from '$lib/client/types';
	import { AlertCircleIcon } from 'lucide-svelte';

	export let invitation: Invitation;
	export let organization: Organization;

	let isOpen = false;

	async function revokeInvitation() {
		await api().Invitation.delete(invitation.organizationId, invitation.id);
		goto(`/org/${organization.slug}/settings/members`);
	}
</script>

<section class="rounded-xl p-4 md:p-8 bg-black/20 text-gray-200 shadow mt-8">
	<h3 class="flex items-center justify-between text-lg font-medium">
		Revoke invitation!
		<AlertCircleIcon class="w-5 h-5" />
	</h3>
	<p class="text-sm mt-2">
		Revoking the invitation will make the invitation email invalid. The user will no longer be able
		to join the organization using the link.
	</p>
	<button class="btn btn-sm mt-4" on:click={() => (isOpen = true)}> Revoke invitation </button>
	<ConfirmModal bind:isOpen action={revokeInvitation} submitText="Revoke invitation">
		<AlertCircleIcon class="w-8 h-8" />
		<p class="py-4 text-center">Are you sure you want to revoke the invitation?</p>
	</ConfirmModal>
</section>
