<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { page } from '$app/stores';
	import type { ActionResult } from '@sveltejs/kit';
	import { Alert, Button, Heading, Input, Label, P, Textarea } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';
	import DeleteOrganizationModal from './_components/DeleteOrganizationModal.svelte';
	import LeaveOrganizationModal from './_components/LeaveOrganizationModal.svelte';

	let isLeaveModalOpen = false;
	let isDeleteModalOpen = false;

	interface Fields {
		name: string;
		description: string;
	}

	let fields: Fields = {
		name: $page.data.organization.name || '',
		description: $page.data.organization.description || ''
	};

	function resetFields() {
		fields = {
			name: $page.data.organization.name || '',
			description: $page.data.organization.description || ''
		};
	}

	async function updateOrganization() {
		const data = new FormData();
		for (const [key, value] of Object.entries(fields)) {
			data.set(key, value);
		}
		const action = '?/updateOrganization';
		const response = await fetch(action, { method: 'post', body: data });
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
		resetFields();
	}
</script>

<Heading customSize="text-xl font-semibold">Organization details</Heading>

<P class="mt-6">
	You can update your organization's information here to ensure accuracy and relevance which will be
	used by AI features.
</P>

<form on:submit|self|stopPropagation|preventDefault={updateOrganization}>
	<div class="mt-6">
		<Label for="name" class="block mb-2">Name:</Label>
		<Input id="name" required bind:value={fields.name} />
	</div>

	<div class="mt-6">
		<Label for="description" class="block mb-2">Description:</Label>
		<Textarea id="description" required bind:value={fields.description} />
	</div>

	<Alert color="red" class="mt-6">
		<div class="flex items-center gap-3">
			<ExclamationCircleOutline slot="icon" class="w-4 h-4" />
			<span class="font-medium">Leave the organization</span>
		</div>
		<p class="mt-2 mb-2 text-sm">
			Exiting this organization will revoke your access to its feedback data and associated
			resources. You will no longer be able to view or manage feedback within this organization. If
			you are the last member in this organization, the organization and all its feedback data will
			be deleted permanently. This action cannot be reversed.
		</p>
		<div class="flex gap-2">
			<Button size="xs" color="red" outline on:click={() => (isLeaveModalOpen = true)}>
				Leave "{$page.data.organization.name}"
			</Button>
		</div>
	</Alert>

	<Alert color="red" class="mt-6">
		<div class="flex items-center gap-3">
			<ExclamationCircleOutline slot="icon" class="w-4 h-4" />
			<span class="font-medium">Delete the organization</span>
		</div>
		<p class="mt-2 mb-4 text-sm">
			Deleting this organization will permanently erase all collected feedback data, including
			associated records and information. Proceeding with this action cannot be undone. Please
			ensure that you have backed up any critical data or exported necessary information before
			deleting the organization. This action cannot be reversed.
		</p>
		<div class="flex gap-2">
			<Button size="xs" color="red" outline on:click={() => (isDeleteModalOpen = true)}>
				Delete "{$page.data.organization.name}"
			</Button>
		</div>
	</Alert>

	<div class="flex justify-end mt-6">
		<Button type="submit" size="sm" color="primary">Update</Button>
	</div>
</form>

<LeaveOrganizationModal organization={$page.data.organization} bind:isOpen={isLeaveModalOpen} />
<DeleteOrganizationModal organization={$page.data.organization} bind:isOpen={isDeleteModalOpen} />
