<script lang="ts">
	import { MinimalLineChart, ValueDonutChart } from '$lib/client/charts';
	import type { TimeSeries } from '$lib/types';
	import { Card, P } from 'flowbite-svelte';

	export let values: TimeSeries<number> = [];

	$: points = values.map((v) => ({ x: v.date, y: Math.round(v.value * 100) }));
	$: value = Math.round(values.reduce((acc, val) => acc + val.value * 100, 0) / values.length);
</script>

<Card shadow={false} padding="none" class="flex flex-1 py-2 px-3">
	<P class="text-sm font-semibold text-gray-700">Customer Satisfaction</P>
	{#if values.length === 0}
		<P class="text-sm text-gray-500">No data</P>
	{:else}
		<main class="flex h-24 gap-2 py-2">
			<div class="flex py-2">
				<ValueDonutChart {value} color="#AB47BC" />
			</div>
			<div class="flex flex-1 border-b border-l border-dashed border-gray-300">
				<MinimalLineChart min={0} max={100} stepX={50} {points} label="CSAT" color="#AB47BC" />
			</div>
		</main>
	{/if}
</Card>
