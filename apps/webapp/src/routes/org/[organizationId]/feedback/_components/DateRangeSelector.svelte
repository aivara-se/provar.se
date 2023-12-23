<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { Button, Dropdown, DropdownItem } from 'flowbite-svelte';
	import { CalendarMonthOutline } from 'flowbite-svelte-icons';

	$: value = $page.url.searchParams.get('range') || '30d';
	$: items = [
		{ name: 'Last 7 days', value: '7d' },
		{ name: 'Last 30 days', value: '30d' }
	].map((item) => ({
		...item,
		active: item.value === value
	}));

	function udpateDateRange(range: string) {
		value = range;
		$page.url.searchParams.set('range', range);
		goto($page.url.toString());
	}
</script>

<div>
	<Button color="light" size="xs">
		<span class="mr-2">{items.find((item) => item.active)?.name}</span>
		<CalendarMonthOutline class="w-3 h-3" />
	</Button>
	<Dropdown>
		{#each items as item}
			<DropdownItem on:click={() => udpateDateRange(item.value)} active={item.active}>
				{item.name}
			</DropdownItem>
		{/each}
	</Dropdown>
</div>
