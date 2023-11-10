<script lang="ts">
	import { page } from '$app/stores';
	import { CenteredLayout, HorizontalSeparator } from '$lib/ui';
	import { Badge, Button, Listgroup, ListgroupItem } from 'flowbite-svelte';
	import { LockOutline, UsersGroupOutline } from 'flowbite-svelte-icons';
</script>

<CenteredLayout>
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

	<HorizontalSeparator text="or" />

	<Button href="/org/create" outline size="xs">
		<UsersGroupOutline class="w-3 h-3 mr-1" />
		Create an organization
	</Button>
</CenteredLayout>

<style>
	.org-item {
		display: flex;
		flex: 1;
		justify-content: space-between;
	}
</style>
