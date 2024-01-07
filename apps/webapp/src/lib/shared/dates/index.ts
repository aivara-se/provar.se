import type { DateRange, FormattedDateRange } from '$lib/types';
import { endOfDay, formatISO, parseISO, subDays } from 'date-fns';

export function parseDateRange(range: FormattedDateRange): DateRange {
	return {
		from: parseISO(range.from + 'T00:00:00.000Z'),
		to: parseISO(range.to + 'T23:59:59.999Z')
	};
}

export function formatDateRangeCompact(range: DateRange): FormattedDateRange {
	return {
		from: formatISO(range.from, { format: 'basic', representation: 'date' }),
		to: formatISO(range.to, { format: 'basic', representation: 'date' })
	};
}

export function formatDateRangeReadable(range: DateRange): string {
	return `${range.from.toLocaleDateString()} - ${range.to.toLocaleDateString()}`;
}

export function createPresetRange(days: number): FormattedDateRange {
	const now = endOfDay(new Date());
	return formatDateRangeCompact({ from: subDays(now, days), to: now });
}
