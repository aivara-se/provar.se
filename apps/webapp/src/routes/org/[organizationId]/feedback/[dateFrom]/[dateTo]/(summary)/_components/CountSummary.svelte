<script lang="ts">
	import { MinimalLineChart } from '$lib/client/charts';
	import type { TimeSeries } from '$lib/types';
	import { Card, P } from 'flowbite-svelte';

	export let values: TimeSeries<number> = [];

	$: points = values.map((v) => ({ x: v.date, y: v.value }));
</script>

<Card shadow={false} padding="none" class="flex flex-1 py-2 px-3">
	<P class="text-sm font-semibold text-gray-700">Daily Feedback Count</P>
	<main class="flex h-24 gap-2 py-2">
		<div class="flex flex-1 border-b border-dashed border-gray-300">
			<MinimalLineChart {points} color="#78909C" />
		</div>
	</main>
</Card>
