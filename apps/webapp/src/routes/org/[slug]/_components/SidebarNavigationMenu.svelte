<script lang="ts">
	interface NagivationItem {
		name: string;
		href: string;
		icon: any;
		active: boolean;
	}

	export let items: NagivationItem[] = [];
</script>

<div class="wrapper">
	{#each items as item}
		<a
			class="text-sm font-medium rounded-lg navigation-item"
			class:active={item.active}
			href={item.href}
		>
			<svelte:component this={item.icon} class="w-3 h-3 mr-4" />
			{item.name}
		</a>
	{/each}
</div>

<style>
	.navigation-item {
		display: flex;
		align-items: center;
		justify-content: flex-start;
		padding: 7px 20px;
		border: solid 1px transparent;
		box-sizing: border-box;
		color: #000;
	}

	.navigation-item.active {
		background: #fff;
		border: solid 1px #ddd;
	}

	.navigation-item + .navigation-item {
		margin-top: 10px;
	}
</style>
