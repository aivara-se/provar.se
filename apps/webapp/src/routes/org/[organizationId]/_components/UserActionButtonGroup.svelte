<script lang="ts">
	import { Button, ButtonGroup } from 'flowbite-svelte';

	interface NagivationItem {
		name: string;
		href?: string;
		onClick?: () => void;
		// TODO: use strict typing when we know which type to use here
		// eslint-disable-next-line @typescript-eslint/no-explicit-any
		icon: any;
		active: boolean;
	}

	export let items: NagivationItem[] = [];
</script>

<ButtonGroup class="mt-2">
	{#each items as item}
		<Button href={item.href} on:click={item.onClick}>
			<span class:text-primary-500={item.active}>
				<svelte:component this={item.icon} class="w-3 h-3 mr-1" />
			</span>
		</Button>
	{/each}
</ButtonGroup>
