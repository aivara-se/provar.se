<script lang="ts">
	import { Badge, Button, Dropdown, DropdownItem } from 'flowbite-svelte';
	import { ChevronDownSolid, LockOutline } from 'flowbite-svelte-icons';

	interface OrganizationInfo {
		id: string;
		slug: string;
		name: string;
		prod: boolean;
	}

	export let value: string;
	export let items: OrganizationInfo[];

	$: organization = items.find((org) => org.slug === value);
</script>

<Button color="light" size="sm" class="dropdown">
	<span class="organization-name">{organization?.name}</span>
	<ChevronDownSolid size="xs" class="w-2 h-2 ml-2" />
</Button>

<Dropdown>
	{#each items as item}
		<DropdownItem>
			<a href={`/org/${item.slug}`} class="org-item">
				<span class="organization-name">{item.name}</span>
				{#if item.prod }
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
		min-width: 120px;
	}

	.organization-name {
		text-overflow: ellipsis;
		overflow: hidden;
		white-space: nowrap;
	}
</style>
