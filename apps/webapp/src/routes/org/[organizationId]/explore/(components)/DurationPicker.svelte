<script lang="ts">
	import { endOfToday, startOfDay, endOfDay, format } from 'date-fns';
	import { CalendarRangeIcon } from 'lucide-svelte';

	export let endDate = endOfToday();
	export let begDate = withOffset(endDate, -30);

	let dateRangePresets = [
		{ label: '1 Day', days: 1, isActive: false },
		{ label: '1 Week', days: 7, isActive: false },
		{ label: '1 Month', days: 30, isActive: false }
	];

	$: {
		const ONE_DAY = 24 * 60 * 60 * 1000;
		const current = Math.ceil((endDate.getTime() - begDate.getTime()) / ONE_DAY);
		const endsToday = endDate.getTime() === endOfToday().getTime();
		for (const preset of dateRangePresets) {
			preset.isActive = endsToday && preset.days === current;
		}
		dateRangePresets = dateRangePresets;
	}

	function withOffset(base: Date, days: number) {
		const ts = base.getTime() + days * 24 * 60 * 60 * 1000;
		return new Date(ts);
	}

	function setDuration(days: number) {
		endDate = endOfToday();
		begDate = withOffset(endDate, -1 * days);
	}
</script>

<div class="dropdown dropdown-end">
	<div tabindex="0" role="button" class="btn btn-ghost btn-sm">
		{format(begDate, 'MMM dd')} - {format(endDate, 'MMM dd')}
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
				value={format(begDate, 'yyyy-MM-dd')}
				on:change={(e) => (begDate = startOfDay(e.currentTarget.value))}
			/>
			<input
				type="date"
				class="input input-bordered input-sm"
				value={format(endDate, 'yyyy-MM-dd')}
				on:change={(e) => (endDate = endOfDay(e.currentTarget.value))}
			/>
		</div>
	</div>
</div>
