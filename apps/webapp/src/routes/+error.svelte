<script lang="ts">
	import { page } from '$app/stores';
	import { Section, Page404 } from 'flowbite-svelte-blocks';
	import { Button } from 'flowbite-svelte';
</script>

<main>
	{#if $page.error?.message === 'Not Found'}
		<Section name="page404">
			<Page404>
				<svelte:fragment slot="h1">404</svelte:fragment>
				<svelte:fragment slot="paragraph">
					<p
						class="mb-4 text-3xl tracking-tight font-bold text-gray-900 md:text-4xl dark:text-white"
					>
						Page not found.
					</p>
					<p class="mb-4 text-lg font-light text-gray-500 dark:text-gray-400">
						Sorry, we can't find that page. Please let us know if you think something is broken.
					</p>
					<Button href="/" size="sm" color="light">Back to Homepage</Button>
				</svelte:fragment>
			</Page404>
		</Section>
	{:else}
		<slot />
	{/if}
</main>

<style>
	main {
		display: flex;
		align-items: center;
		justify-content: center;
		min-height: 100vh;
	}
</style>
