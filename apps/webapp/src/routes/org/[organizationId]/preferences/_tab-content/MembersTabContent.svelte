<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { selectedOrg } from '$lib/stores/selected-org';
	import type { ActionResult } from '@sveltejs/kit';
	import { Alert, Button, Heading, Modal, P } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';

	let name = $selectedOrg?.name || '';

	let isLeaveModalOpen = false;

	async function leaveOrganization() {
		const data = new FormData();
		const response = await fetch('?/leaveOrganization', {
			method: 'post',
			body: data
		});
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
	}
</script>

<Heading customSize="text-xl font-semibold">Organization members</Heading>

<P class="mt-6">
	Below is the list of individuals currently associated with this organization. You can view and
	manage members here.
</P>

<Alert color="red" class="mt-6">
	<div class="flex items-center gap-3">
		<ExclamationCircleOutline slot="icon" class="w-4 h-4" />
		<span class="font-medium">Leave the organization</span>
	</div>
	<p class="mt-2 mb-2 text-sm">
		Exiting this organization will revoke your access to its feedback data and associated resources.
		You will no longer be able to view or manage feedback within this organization.
	</p>
	<p class="mt-2 mb-4 text-sm">
		<strong>Please note:</strong> If you are the last member in this organization, the organization and
		all its feedback data will be deleted permanently. This action cannot be reversed.
	</p>
	<div class="flex gap-2">
		<Button size="xs" color="red" outline on:click={() => (isLeaveModalOpen = true)}>
			{#if $selectedOrg?.members.length === 1}
				Delete "{name}"
			{:else}
				Leave "{name}"
			{/if}
		</Button>
	</div>
</Alert>

<Modal bind:open={isLeaveModalOpen} size="sm" autoclose>
	<div class="text-center">
		<ExclamationCircleOutline class="mx-auto mb-4 text-gray-500 w-8 h-8" />
		<h3 class="mb-5 text-lg font-normal text-gray-800">
			{#if $selectedOrg?.members.length === 1}
				Are you sure you want to delete "{name}"?
			{:else}
				Are you sure you want to leave "{name}"?
			{/if}
		</h3>
		<Button color="red" class="me-2" on:click={leaveOrganization}>Yes, I'm sure</Button>
		<Button color="alternative">No, cancel</Button>
	</div>
</Modal>
