<script lang="ts">
	import { AlertCircleIcon } from 'lucide-svelte';
	import { ConfirmModal } from '$lib/client/forms';
	import type { Organization } from '$lib/types';

	export let organization: Organization;

	let isOpen = false;
</script>

<section class="mt-8 rounded-lg p-4 bg-gray-100 text-gray-900 dark:bg-gray-950 dark:text-gray-200">
	<h3 class="flex items-center justify-between text-lg font-medium">Create an API Key</h3>
	<p class="text-sm">
		API keys are used to authenticate your application with the API. These keys allow your
		application to upload feedback data.
	</p>
	<button class="btn btn-sm btn-neutral mt-4" on:click={() => (isOpen = true)}>
		Create an API Key
	</button>
	<ConfirmModal bind:isOpen submitPath="?/deleteOrganization" submitText="Create an API Key">
		<AlertCircleIcon class="w-8 h-8" />
		<p class="py-4 text-center">
			Are you sure you want to delete {organization?.name}?
		</p>
	</ConfirmModal>
</section>
