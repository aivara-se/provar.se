<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/client/api';
	import { ConfirmModal } from '$lib/client/forms';
	import type { Organization } from '$lib/client/types';
	import { AlertCircleIcon } from 'lucide-svelte';

	export let organization: Organization;

	let isOpen = false;

	async function deleteOrganization() {
		await api().Organization.delete(organization.id);
		goto('/org');
	}
</script>

<section class="rounded-xl p-4 md:p-8 bg-black/20 text-gray-200 shadow mt-8">
	<h3 class="flex items-center justify-between text-lg font-medium">
		Delete the organization!
		<AlertCircleIcon class="w-5 h-5" />
	</h3>
	<p class="text-sm mt-2">
		Deleting this organization will permanently erase all collected feedback data, including
		associated records and information. This action cannot be reversed.
	</p>
	<button class="btn btn-sm mt-4" on:click={() => (isOpen = true)}>
		Delete {organization?.name}
	</button>
	<ConfirmModal bind:isOpen action={deleteOrganization} submitText="Delete {organization?.name}">
		<AlertCircleIcon class="w-8 h-8" />
		<p class="py-4 text-center">
			Are you sure you want to delete {organization?.name}?
		</p>
	</ConfirmModal>
</section>
