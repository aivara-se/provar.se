<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/client/api';
	import { ConfirmModal } from '$lib/client/forms';
	import type { Organization, User } from '$lib/client/types';
	import { AlertCircleIcon } from 'lucide-svelte';

	export let organization: Organization;
	export let user: User;

	let isOpen = false;

	async function leaveOrganization() {
		await api().Organization.removeMember(organization.id, user.id);
		goto('/org');
	}
</script>

<section class="mt-8 rounded-lg p-4 bg-gray-100 text-gray-900 dark:bg-gray-950 dark:text-gray-200">
	<h3 class="flex items-center justify-between text-lg font-medium">
		Leave the organization!
		<AlertCircleIcon class="w-5 h-5" />
	</h3>
	<p class="text-sm mt-2">
		Exiting this organization will revoke your access to its feedback data and associated resources.
		You will no longer be able to view or manage feedback within this organization.
	</p>
	<button class="btn btn-sm btn-neutral mt-4" on:click={() => (isOpen = true)}>
		Leave {organization?.name}
	</button>
	<ConfirmModal bind:isOpen action={leaveOrganization} submitText="Leave {organization?.name}">
		<AlertCircleIcon class="w-8 h-8" />
		<p class="py-4 text-center">
			Are you sure you want to leave {organization?.name}?
		</p>
	</ConfirmModal>
</section>
