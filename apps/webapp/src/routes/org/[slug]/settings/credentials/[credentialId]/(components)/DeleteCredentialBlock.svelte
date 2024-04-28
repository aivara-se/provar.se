<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/client/api';
	import { ConfirmModal } from '$lib/client/forms';
	import type { Credential, Organization } from '$lib/client/types';
	import { AlertCircleIcon } from 'lucide-svelte';

	export let credential: Credential;
	export let organization: Organization;

	let isOpen = false;

	async function revokeCredential() {
		await api().Credential.delete(credential.organizationId, credential.id);
		goto(`/org/${organization.slug}/settings/credentials`);
	}
</script>

<section class="rounded-xl p-4 md:p-8 bg-black/20 text-gray-200 shadow mt-8">
	<h3 class="flex items-center justify-between text-lg font-medium">
		Delete the credential!
		<AlertCircleIcon class="w-5 h-5" />
	</h3>
	<p class="text-sm mt-2">
		Deleting this credential will remove access from applications using this credential. Future API
		requests using this credential will fail. This action cannot be reversed.
	</p>
	<button class="btn btn-sm mt-4" on:click={() => (isOpen = true)}>
		Delete {credential?.name}
	</button>
	<ConfirmModal bind:isOpen action={revokeCredential} submitText="Delete {credential?.name}">
		<AlertCircleIcon class="w-8 h-8" />
		<p class="py-4 text-center">
			Are you sure you want to delete {credential?.name}?
		</p>
	</ConfirmModal>
</section>
