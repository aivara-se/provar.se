<script lang="ts">
	import { AlertCircleIcon } from 'lucide-svelte';
	import { ConfirmModal } from '$lib/client/forms';
	import type { Project } from '$lib/types';

	export let project: Project;

	let isOpen = false;
</script>

<section class="mt-8 rounded-lg p-4 bg-gray-100 text-gray-900 dark:bg-gray-950 dark:text-gray-200">
	<h3 class="flex items-center justify-between text-lg font-medium">
		Delete the project!
		<AlertCircleIcon class="w-5 h-5" />
	</h3>
	<p class="text-sm mt-2">
		Deleting this project will remove the project and unlink feedback associated with this project.
		Feedback data is not removed. This action cannot be reversed.
	</p>
	<button class="btn btn-sm btn-neutral mt-4" on:click={() => (isOpen = true)}>
		Delete {project?.name}
	</button>
	<ConfirmModal bind:isOpen submitPath="?/removeProject" submitText="Delete {project?.name}">
		<AlertCircleIcon class="w-8 h-8" />
		<p class="py-4 text-center">
			Are you sure you want to delete {project?.name}?
		</p>
	</ConfirmModal>
</section>
