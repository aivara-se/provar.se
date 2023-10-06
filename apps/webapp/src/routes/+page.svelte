<script lang="ts">
	import { page } from '$app/stores';
	import { Badge, Button, Listgroup, ListgroupItem } from 'flowbite-svelte';
	import { LockOutline, UsersGroupOutline } from 'flowbite-svelte-icons';
</script>

<div class="wrapper">
	<main>
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

		<div class="separator">
			<span class="bg-white px-2 text-sm text-gray-500">or</span>
		</div>

		<Button href="/org/create" outline size="xs">
			<UsersGroupOutline class="w-3 h-3 mr-1" />
			Create an organization
		</Button>
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

	.separator {
		display: flex;
		align-items: center;
		justify-content: center;
		margin: 25px 0;
		width: 160px;
		height: 1px;
		background: #e0e0e0;
	}
</style>
