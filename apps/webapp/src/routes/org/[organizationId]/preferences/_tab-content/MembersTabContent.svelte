<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { selectedOrg } from '$lib/stores/selected-org';
	import type { ActionResult } from '@sveltejs/kit';
	import { Alert, Button, Heading, P } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';

	let name = $selectedOrg?.name || '';

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
	<p class="mt-2 mb-4 text-sm">
		Exiting this organization will revoke your access to its feedback data and associated resources.
		You will no longer be able to view or manage feedback within this organization. Please note:
		Upon your departure, if you are the last member, the organization and all its feedback data will
		be deleted permanently. Ensure you've saved or transferred any essential information before
		leaving, as this action cannot be reversed.
	</p>
	<div class="flex gap-2">
		<Button size="xs" color="red" outline on:click={leaveOrganization}>Leave "{name}"</Button>
	</div>
</Alert>
