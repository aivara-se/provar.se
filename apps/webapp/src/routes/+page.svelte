<script lang="ts">
	import { page } from '$app/stores';
	import { Badge, Heading, Listgroup, ListgroupItem } from 'flowbite-svelte';
	import { LockOutline } from 'flowbite-svelte-icons';
</script>

<div class="wrapper">
	<main>
		<Heading tag="h1" customSize="text-2xl" class="mb-4">Organizations</Heading>

		<Listgroup active class="w-48">
			{#each $page.data.organizations as org}
				<ListgroupItem>
					<a href={`/org/${org.id}`} class="org-item">
						{org.name}
						{#if org.environment === 'production'}
							<Badge class="text-xs font-semibold ml-2">
								<LockOutline class="w-3 h-3 mr-1" />
							</Badge>
						{/if}
					</a>
				</ListgroupItem>
			{/each}
		</Listgroup>
	</main>
</div>

<style>
	.wrapper {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		height: 100vh;
	}

	main {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		text-align: center;
	}

	.org-item {
		display: flex;
		flex: 1;
		justify-content: space-between;
	}
</style>
