<script lang="ts">
	import Chart from 'chart.js/auto';
	import 'chartjs-adapter-date-fns';
	import { subDays } from 'date-fns';
	import { enUS } from 'date-fns/locale';
	import { onMount } from 'svelte';

	export let range: string = '30d';
	export let label: string = 'Count';
	export let points: { x: string; y: number }[] = [];
	export let color: string = '#4CAF50';

	export let min: number | undefined = undefined;
	export let max: number | undefined = undefined;
	export let stepX: number | undefined = undefined;

	let element: HTMLCanvasElement;

	function getLabelsFromRange(range: string) {
		const [amount] = range.split(/(?<=\d)(?=[a-z])/);
		const labels = [];
		let date = new Date();
		for (let i = 0; i < parseInt(amount); i++) {
			labels.push(date.toISOString());
			date = subDays(date, 1);
		}
		return labels.reverse();
	}

	$: labels = getLabelsFromRange(range);

	onMount(() => {
		new Chart(element, {
			type: 'bar',
			data: {
				labels: labels,
				datasets: [
					{
						data: points,
						borderColor: color,
						backgroundColor: color
					}
				]
			},
			options: {
				animation: false,
				responsive: true,
				maintainAspectRatio: false,
				scales: {
					x: {
						display: false,
						type: 'time',
						time: { unit: 'day', minUnit: points.length > 50 ? 'month' : 'day' },
						adapters: { date: { locale: enUS } }
					},
					y: {
						display: false,
						beginAtZero: true,
						min,
						max,
						...(stepX ? { ticks: { stepSize: stepX } } : {})
					}
				},
				plugins: {
					legend: { display: false },
					tooltip: {
						displayColors: false,
						callbacks: {
							title: (context) => {
								const date = new Date(context[0].parsed.x);
								return date.toLocaleString([], { month: 'short', day: 'numeric' });
							},
							label: (context) => {
								const value = context.parsed.y;
								return `${label}: ${value}`;
							}
						}
					},
					decimation: {
						enabled: true,
						algorithm: 'min-max'
					}
				}
			}
		});
	});
</script>

<canvas bind:this={element} />
