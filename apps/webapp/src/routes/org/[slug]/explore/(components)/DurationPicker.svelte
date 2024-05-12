<script lang="ts">
	import { goto } from '$app/navigation';
	import { formatDateRange } from '$lib/client/dates';
	import { type DateRange } from '$lib/client/types';
	import { addDays, endOfDay, endOfToday, format, startOfDay } from 'date-fns';
	import { CalendarRangeIcon } from 'lucide-svelte';

	const ONE_DAY = 24 * 60 * 60 * 1000;

	export let range: Partial<DateRange> = {};

	$: rangePresets = [
		{ label: 'Today', days: 1, isActive: isActiveRange(1) },
		{ label: '30 Days', days: 30, isActive: isActiveRange(30) }
	];

	function isActiveRange(days: number) {
		if (!range.from || !range.to) {
			return false;
		}
		const rangeDays = Math.ceil((range.to.getTime() - range.from.getTime()) / ONE_DAY);
		const endsToday = Boolean(range.to.getTime() === endOfToday().getTime());
		return endsToday && days === rangeDays;
	}

	function setDuration(days: number) {
		const newEndDate = endOfToday();
		const newBegDate = startOfDay(addDays(newEndDate, 1 + -1 * days));
		setDateRange(newBegDate, newEndDate);
	}

	function setDateRange(from?: Date, to?: Date) {
		const value = formatDateRange({ from, to });
		const url = new URL(window.location.href);
		if (value.from) {
			url.searchParams.set('from', value.from);
		} else {
			url.searchParams.delete('from');
		}
		if (value.to) {
			url.searchParams.set('to', value.to);
		} else {
			url.searchParams.delete('to');
		}
		goto(url.toString());
	}

	function stringifyDateRange(from?: Date, to?: Date) {
		const fromString = from ? format(from, 'MMM dd') : '';
		const toString = to ? format(to, 'MMM dd') : '';
		if (!fromString && !toString) {
			return 'Set Duration';
		}
		if (fromString === toString) {
			return fromString;
		}
		if (!fromString && toString) {
			return `Up to ${toString}`;
		}
		if (fromString && !toString) {
			return `From ${fromString}`;
		}
		return `${fromString} - ${toString}`;
	}
</script>

<div class="dropdown dropdown-end">
	<div tabindex="0" role="button" class="btn btn-ghost btn-sm">
		{stringifyDateRange(range.from, range.to)}
		<CalendarRangeIcon class="w-3.5 h-3.5 ml-1" />
	</div>
	<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
	<div
		tabindex="0"
		class="dropdown-content z-[1] card card-compact w-56 mt-1 p-0 bg-base-100 shadow-lg"
	>
		<div class="card-body">
			<div class="flex flex-row justify-center gap-2">
				{#each rangePresets as preset}
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
				value={range.from ? format(range.from, 'yyyy-MM-dd') : ''}
				on:change={(e) =>
					setDateRange(
						e.currentTarget.value ? startOfDay(e.currentTarget.value) : undefined,
						range.to
					)}
			/>
			<input
				type="date"
				class="input input-bordered input-sm"
				value={range.to ? format(range.to, 'yyyy-MM-dd') : ''}
				on:change={(e) =>
					setDateRange(
						range.from,
						e.currentTarget.value ? endOfDay(e.currentTarget.value) : undefined
					)}
			/>
		</div>
	</div>
</div>
