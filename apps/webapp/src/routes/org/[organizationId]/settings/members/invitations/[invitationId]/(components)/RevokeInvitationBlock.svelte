<script lang="ts">
	import { AlertCircleIcon } from 'lucide-svelte';
	import { ConfirmModal } from '$lib/client/forms';
	import type { Invitation } from '$lib/types';

	export let invitation: Invitation;

	let isOpen = false;
</script>

<section class="mt-8 rounded-lg p-4 bg-gray-100 text-gray-900 dark:bg-gray-950 dark:text-gray-200">
	<h3 class="flex items-center justify-between text-lg font-medium">
		Revoke invitation!
		<AlertCircleIcon class="w-5 h-5" />
	</h3>
	<p class="text-sm mt-2">
		Revoking the invitation will make the invitation email invalid. The user will no longer be able
		to join the organization using the link.
	</p>
	<button class="btn btn-sm btn-neutral mt-4" on:click={() => (isOpen = true)}>
		Revoke invitation
	</button>
	<ConfirmModal bind:isOpen submitPath="?/revokeInvitation" submitText="Revoke invitation">
		<AlertCircleIcon class="w-8 h-8" />
		<p class="py-4 text-center">Are you sure you want to revoke the invitation?</p>
	</ConfirmModal>
</section>
