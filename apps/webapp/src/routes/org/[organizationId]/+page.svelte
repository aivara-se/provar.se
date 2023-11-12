<script lang="ts">
	import { page } from '$app/stores';
	import GettingStartedGuide from './_page-content/GettingStartedGuide.svelte';

	/**
	 * Check if the user has credentials when the page is loaded. During the onboarding
	 * process, the user will create credentials and we should not hide the onboarding
	 * guide when that happens. Because of that `hasCredentials` is only set on page load.
	 */
	let hasCredentials = $page.data.credentials && $page.data.credentials.length > 0;
</script>

{#if !hasCredentials}
	<GettingStartedGuide />
{/if}
