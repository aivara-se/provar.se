<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { createPresetRange, formatDateRangeReadable, parseDateRange } from '$lib/shared/dates';
	import type { FormattedDateRange } from '$lib/types';
	import { Button, Dropdown, DropdownItem } from 'flowbite-svelte';
	import { CalendarMonthOutline } from 'flowbite-svelte-icons';

	$: value = { from: $page.params.dateFrom, to: $page.params.dateTo };

	$: items = [
		{ name: 'Last week', value: createPresetRange(7) },
		{ name: 'Last month', value: createPresetRange(30) },
		{ name: 'Last year', value: createPresetRange(365) }
	].map((item) => ({ ...item, active: areRangesEqual(item.value, value) }));

	function areRangesEqual(x: FormattedDateRange, y: FormattedDateRange): boolean {
		return x.from === y.from && x.to === y.to;
	}

	function udpateDateRange(range: FormattedDateRange) {
		const url = new URL($page.url.toString());
		url.searchParams.set('page', '1');
		goto(`/org/${$page.params.organizationId}/feedback/${range.from}/${range.to}` + url.search);
	}
</script>

<div>
	<Button color="light" size="xs">
		<span class="mr-2">{formatDateRangeReadable(parseDateRange(value))}</span>
		<CalendarMonthOutline class="w-3 h-3" />
	</Button>
	<Dropdown>
		{#each items as item}
			<DropdownItem on:click={() => udpateDateRange(item.value)}>
				{item.name}
			</DropdownItem>
		{/each}
	</Dropdown>
</div>
