<script lang="ts">
	import { ChevronRightSolid } from 'flowbite-svelte-icons';

	interface NagivationItem {
		name: string;
		href: string;
		active: boolean;
	}

	export let items: NagivationItem[] = [];
</script>

<div class="wrapper">
	{#each items as item}
		<a class="text-sm font-medium rounded-lg navigation-item" href={item.href}>
			{item.name}
			{#if item.active}
				<ChevronRightSolid class="w-2 h-2" />
			{/if}
		</a>
	{/each}
</div>

<style>
	.navigation-item {
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding: 7px 20px;
		border: solid 1px transparent;
		box-sizing: border-box;
		color: #000;
	}

	.navigation-item:hover {
		background: #fff;
		border: solid 1px #ddd;
	}

	.navigation-item + .navigation-item {
		margin-top: 10px;
	}
</style>
