<script lang="ts">
	import { Badge, Button, Dropdown, DropdownItem } from 'flowbite-svelte';
	import { ChevronDownSolid, LockOutline } from 'flowbite-svelte-icons';

	interface OrganizationInfo {
		id: string;
		name: string;
		environment: string;
	}

	export let organization: OrganizationInfo;
	export let organizations: OrganizationInfo[];
</script>

<Button color="light" size="sm">
	{organization.name}
	<ChevronDownSolid size="xs" class="w-2 h-2 ml-2" />
</Button>

<Dropdown>
	{#each organizations as org}
		<DropdownItem>
			<a href={`/org/${org.id}`} class="org-item">
				{org.name}
				{#if org.environment === 'production'}
					<Badge class="text-xs font-semibold ml-2">
						<LockOutline class="w-3 h-3 mr-1" />
					</Badge>
				{/if}
			</a>
		</DropdownItem>
	{/each}
</Dropdown>

<style>
	.org-item {
		display: flex;
		flex: 1;
		justify-content: space-between;
	}
</style>
