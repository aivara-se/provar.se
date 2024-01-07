<script lang="ts">
	import Chart from 'chart.js/auto';
	import { onMount } from 'svelte';

	export let value: number = 0;
	export let color: string = '#4CAF50';

	let element: HTMLCanvasElement;

	onMount(() => {
		new Chart(element, {
			type: 'doughnut',
			data: {
				datasets: [
					{
						data: [Math.floor(value), 100 - Math.floor(value)],
						backgroundColor: [color, '#EEEEEE']
					}
				]
			},
			options: {
				animation: false,
				cutout: '60%',
				plugins: {
					legend: { display: false },
					tooltip: { enabled: false }
				}
			}
		});
	});
</script>

<div class="flex-1 w-full relative">
	<canvas bind:this={element} />
	<div class="flex items-center justify-center absolute w-full h-full top-0 pointer-events-none">
		<span class="text-xs font-semibold">{Math.floor(value)}</span>
	</div>
</div>
