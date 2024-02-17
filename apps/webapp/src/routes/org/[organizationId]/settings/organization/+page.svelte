<script lang="ts">
	import { ConfirmModal } from '$lib/client/forms';
	import { DashLayout } from '$lib/client/layout';
	import { AlertCircleIcon } from 'lucide-svelte';
	import UpdateOrganizationForm from './(forms)/UpdateOrganizationForm.svelte';
	import route from './route.meta';

	export let data;

	let isLeaveModalOpen = false;
	let isDeleteModalOpen = false;
</script>

<DashLayout {route}>
	<UpdateOrganizationForm formData={data.UpdateOrganizationForm} />

	<section
		class="mt-8 rounded-lg p-4 bg-gray-100 text-gray-900 dark:bg-gray-950 dark:text-gray-200"
	>
		<h3 class="flex items-center justify-between text-lg font-medium">
			Leave the organization!
			<AlertCircleIcon class="w-5 h-5" />
		</h3>
		<p class="text-sm">
			Exiting this organization will revoke your access to its feedback data and associated
			resources. You will no longer be able to view or manage feedback within this organization. If
			you are the last member in this organization, the organization and all its feedback data will
			be deleted permanently. This action cannot be reversed.
		</p>
		<button class="btn btn-sm btn-neutral mt-4" on:click={() => (isLeaveModalOpen = true)}>
			Leave {data.organization?.name}
		</button>
		<ConfirmModal
			bind:isOpen={isLeaveModalOpen}
			submitPath="?/leaveOrganization"
			submitText="Leave {data.organization?.name}"
		>
			<AlertCircleIcon class="w-8 h-8" />
			<p class="py-4 text-center">
				Are you sure you want to leave {data.organization?.name}?
			</p>
		</ConfirmModal>
	</section>

	<section
		class="mt-8 rounded-lg p-4 bg-gray-100 text-gray-900 dark:bg-gray-950 dark:text-gray-200"
	>
		<h3 class="flex items-center justify-between text-lg font-medium">
			Delete the organization!
			<AlertCircleIcon class="w-5 h-5" />
		</h3>
		<p class="text-sm">
			Deleting this organization will permanently erase all collected feedback data, including
			associated records and information. Proceeding with this action cannot be undone. Please
			ensure that you have backed up any critical data or exported necessary information before
			deleting the organization. This action cannot be reversed.
		</p>
		<button class="btn btn-sm btn-neutral mt-4" on:click={() => (isDeleteModalOpen = true)}>
			Delete {data.organization?.name}
		</button>
		<ConfirmModal
			bind:isOpen={isDeleteModalOpen}
			submitPath="?/deleteOrganization"
			submitText="Delete {data.organization?.name}"
		>
			<AlertCircleIcon class="w-8 h-8" />
			<p class="py-4 text-center">
				Are you sure you want to delete {data.organization?.name}?
			</p>
		</ConfirmModal>
	</section>
</DashLayout>
