<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { addDays, endOfDay, endOfToday, format, startOfDay } from 'date-fns';
	import { CalendarRangeIcon } from 'lucide-svelte';
	import { onMount } from 'svelte';

	export let range = {
		beg: startOfDay(addDays(endOfToday(), -29)),
		end: endOfToday()
	};

	let dateRangePresets = [
		{ label: '1 Day', days: 1, isActive: false },
		{ label: '1 Week', days: 7, isActive: false },
		{ label: '1 Month', days: 30, isActive: false }
	];

	onMount(() => {
		const begDateParam = $page.url.searchParams.get('beg');
		const endDateParam = $page.url.searchParams.get('end');
		if (!begDateParam || !endDateParam) {
			setDateRange(range);
		}
	});

	$: {
		const begDateParam = $page.url.searchParams.get('beg');
		const endDateParam = $page.url.searchParams.get('end');
		if (begDateParam && endDateParam) {
			range = {
				beg: startOfDay(new Date(begDateParam)),
				end: endOfDay(new Date(endDateParam))
			};
		}
	}

	$: {
		const ONE_DAY = 24 * 60 * 60 * 1000;
		const current = Math.ceil((range.end.getTime() - range.beg.getTime()) / ONE_DAY);
		const endsToday = range.end.getTime() === endOfToday().getTime();
		for (const preset of dateRangePresets) {
			preset.isActive = endsToday && preset.days === current;
		}
		dateRangePresets = dateRangePresets;
	}

	function setDuration(days: number) {
		const newEndDate = endOfToday();
		const newBegDate = startOfDay(addDays(newEndDate, 1 + -1 * days));
		setDateRange({ beg: newBegDate, end: newEndDate });
	}

	function setDateRange(range: { beg: Date; end: Date }) {
		const url = new URL($page.url.toString());
		url.searchParams.set('beg', format(range.beg, 'yyyy-MM-dd'));
		url.searchParams.set('end', format(range.end, 'yyyy-MM-dd'));
		goto(url.toString());
	}
</script>

<div class="dropdown dropdown-end">
	<div tabindex="0" role="button" class="btn btn-ghost btn-sm">
		{format(range.beg, 'MMM dd')} - {format(range.end, 'MMM dd')}
		<CalendarRangeIcon class="w-3.5 h-3.5 ml-1" />
	</div>
	<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
	<div
		tabindex="0"
		class="dropdown-content z-[1] card card-compact w-56 mt-1 p-0 bg-base-100 shadow-lg"
	>
		<div class="card-body">
			<div class="flex flex-row justify-center gap-2">
				{#each dateRangePresets as preset}
					<button
						class="btn btn-xs {preset.isActive ? 'btn-primary' : 'btn-ghost'}"
						on:click={() => setDuration(preset.days)}
					>
						{preset.label}
					</button>
				{/each}
			</div>
			<input
				type="date"
				class="input input-bordered input-sm mt-1"
				value={format(range.beg, 'yyyy-MM-dd')}
				on:change={(e) => setDateRange({ beg: startOfDay(e.currentTarget.value), end: range.end })}
			/>
			<input
				type="date"
				class="input input-bordered input-sm"
				value={format(range.end, 'yyyy-MM-dd')}
				on:change={(e) => setDateRange({ beg: range.beg, end: endOfDay(e.currentTarget.value) })}
			/>
		</div>
	</div>
</div>
