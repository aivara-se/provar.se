<script>
	import { page } from '$app/stores';
	import { Alert } from 'flowbite-svelte';
	import { InfoCircleSolid } from 'flowbite-svelte-icons';

	$: project = $page.data.project;
</script>

<Alert color="blue" rounded={false} class="border-t-4">
	<InfoCircleSolid slot="icon" class="w-4 h-4" />
	<span class="font-medium">Project details: {project?.name}</span><br />
	This page is under construction.
</Alert>
