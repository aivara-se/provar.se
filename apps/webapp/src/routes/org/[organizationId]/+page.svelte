<script lang="ts">
	import { page } from '$app/stores';
	import GettingStartedGuide from './_page-content/GettingStartedGuide.svelte';

	let hasCredentials = $page.data.credentials.length > 0;
</script>

{#if !hasCredentials}
	<GettingStartedGuide />
{/if}
