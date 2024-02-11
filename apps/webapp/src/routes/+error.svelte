<script lang="ts">
	import { page } from '$app/stores';
	import Page404 from './(errors)/Page404.svelte';
	import Page500 from './(errors)/Page500.svelte';
</script>

{#if $page.error?.message === 'Not Found'}
	<Page404 />
{:else}
	<Page500 />
{/if}
