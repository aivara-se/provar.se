<script lang="ts">
	import { Button, Dropdown, DropdownDivider, DropdownItem, P } from 'flowbite-svelte';
	import { ChevronDownSolid } from 'flowbite-svelte-icons';

	interface OrganizationInfo {
		id: string;
		name: string;
	}

	export let current: string;
	export let items: OrganizationInfo[];

	$: organization = items.find((org) => org.id === current);
</script>

<Button color="light" size="xs">
	<div class="dropdown">
		<span class="organization-name">{organization?.name}</span>
		<ChevronDownSolid size="xs" class="w-2 h-2 ml-2" />
	</div>
</Button>

<Dropdown>
	{#each items as item}
		<DropdownItem>
			<a href={`/org/${item.id}`} class="org-item">
				<span class="organization-name">{item.name}</span>
			</a>
		</DropdownItem>
	{/each}

	<DropdownDivider />

	<DropdownItem>
		<a href={`/org/create`} class="create-organization">
			<P size="xs" color="text-gray-700">+ Create Organization</P>
		</a>
	</DropdownItem>
</Dropdown>

<style>
	.dropdown {
		display: flex;
		flex: 1;
		align-items: center;
		justify-content: space-between;
		width: 122px;
	}

	.org-item {
		display: flex;
		flex: 1;
		justify-content: space-between;
		min-width: 120px;
	}

	.create-organization {
		display: flex;
		flex: 1;
		justify-content: space-between;
		min-width: 120px;
	}

	.organization-name {
		display: inline-block;
		flex: 1;
		text-overflow: ellipsis;
		overflow: hidden;
		white-space: nowrap;
	}
</style>
