import { formatISO, parseISO } from 'date-fns';
import type { DateRange, FormattedDateRange } from '../types';

export function parseDateRange(range: Partial<FormattedDateRange>): Partial<DateRange> {
	const parsed = {} as Partial<DateRange>;
	if (range.from) {
		parsed.from = parseISO(range.from + 'T00:00:00.000Z');
	}
	if (range.to) {
		parsed.to = parseISO(range.to + 'T23:59:59.999Z');
	}
	return parsed;
}

export function formatDateString(date?: string): string {
	if (!date) {
		return '';
	}
	return formatISO(new Date(date), { format: 'basic', representation: 'date' });
}

export function formatDateRange(range: Partial<DateRange>): Partial<FormattedDateRange> {
	const formatted = {} as Partial<FormattedDateRange>;
	if (range.from) {
		formatted.from = formatDateString(range.from.toISOString());
	}
	if (range.to) {
		formatted.to = formatDateString(range.to.toISOString());
	}
	return formatted;
}
