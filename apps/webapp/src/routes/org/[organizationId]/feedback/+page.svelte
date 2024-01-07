<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { createPresetRange } from '$lib/shared/dates';
	import { onMount } from 'svelte';

	onMount(() => {
		const range = createPresetRange(30);
		const url = `/org/${$page.params.organizationId}/feedback/${range.from}/${range.to}`;
		goto(url, { replaceState: true });
	});
</script>
