<script lang="ts">
	import { Listgroup, ListgroupItem } from 'flowbite-svelte';

	interface NagivationItem {
		name: string;
		href: string;
		// TODO: use strict typing when we know which type to use here
		// eslint-disable-next-line @typescript-eslint/no-explicit-any
		icon: any;
		active: boolean;
	}

	export let items: NagivationItem[] = [];
</script>

<Listgroup active>
	{#each items as item}
		<ListgroupItem href={item.href} current={item.active}>
			<div class="navigation-item {!item.active ? 'text-gray-900' : ''}">
				<svelte:component this={item.icon} class="w-3 h-3 mr-4" />
				{item.name}
			</div>
		</ListgroupItem>
	{/each}
</Listgroup>

<style>
	.navigation-item {
		display: flex;
		align-items: center;
	}
</style>
