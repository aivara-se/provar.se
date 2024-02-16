<script lang="ts">
	import { AlertCircleIcon } from 'lucide-svelte';
	import { enhance } from '$app/forms';
	import UpdateOrganizationForm from './(forms)/UpdateOrganizationForm.svelte';
	import { CloseButton } from '$lib/client/forms';

	export let data;

	let isLeaveModalOpen = false;
	let isDeleteModalOpen = false;
</script>

<UpdateOrganizationForm formData={data.UpdateOrganizationForm} />

<section class="mt-8 rounded-lg p-4 bg-gray-100 text-gray-900 dark:bg-gray-950 dark:text-gray-200">
	<h3 class="flex items-center justify-between text-lg font-medium">
		Leave the organization!
		<AlertCircleIcon class="w-5 h-5" />
	</h3>
	<p class="text-sm">
		Exiting this organization will revoke your access to its feedback data and associated resources.
		You will no longer be able to view or manage feedback within this organization. If you are the
		last member in this organization, the organization and all its feedback data will be deleted
		permanently. This action cannot be reversed.
	</p>
	<button class="btn btn-sm btn-neutral mt-4" on:click={() => (isLeaveModalOpen = true)}>
		Leave {data.organization?.name}
	</button>
	<dialog class="modal" class:modal-open={isLeaveModalOpen}>
		<div class="modal-box flex flex-col items-center">
			<CloseButton on:click={() => (isLeaveModalOpen = false)} />
			<AlertCircleIcon class="w-8 h-8" />
			<p class="py-4 text-center">Are you sure you want to leave {data.organization?.name}?</p>
			<footer class="flex justify-around gap-4">
				<form method="POST" action="?/leaveOrganization" use:enhance>
					<button class="btn btn-neutral">Leave {data.organization?.name}</button>
				</form>
				<button class="btn btn-ghost" on:click={() => (isLeaveModalOpen = false)}> Cancel </button>
			</footer>
		</div>
	</dialog>
</section>

<section class="mt-8 rounded-lg p-4 bg-gray-100 text-gray-900 dark:bg-gray-950 dark:text-gray-200">
	<h3 class="flex items-center justify-between text-lg font-medium">
		Delete the organization!
		<AlertCircleIcon class="w-5 h-5" />
	</h3>
	<p class="text-sm">
		Deleting this organization will permanently erase all collected feedback data, including
		associated records and information. Proceeding with this action cannot be undone. Please ensure
		that you have backed up any critical data or exported necessary information before deleting the
		organization. This action cannot be reversed.
	</p>
	<button class="btn btn-sm btn-neutral mt-4" on:click={() => (isDeleteModalOpen = true)}>
		Delete {data.organization?.name}
	</button>
	<dialog class="modal" class:modal-open={isDeleteModalOpen}>
		<div class="modal-box flex flex-col items-center">
			<CloseButton on:click={() => (isDeleteModalOpen = false)} />
			<AlertCircleIcon class="w-8 h-8" />
			<p class="py-4 text-center">Are you sure you want to delete {data.organization?.name}?</p>
			<footer class="flex justify-around gap-4">
				<form method="POST" action="?/deleteOrganization" use:enhance>
					<button class="btn btn-neutral">Delete {data.organization?.name}</button>
				</form>
				<button class="btn btn-ghost" on:click={() => (isDeleteModalOpen = false)}> Cancel </button>
			</footer>
		</div>
	</dialog>
</section>
