export type DatePeriod = [Date, Date];

export function getDefaultPeriod(): DatePeriod {
	const startDate = new Date();
	startDate.setFullYear(startDate.getFullYear() - 1);
	const endDate = new Date();
	return [startDate, endDate];
}

export function randomDate(period: DatePeriod): Date {
	const [start, end] = period;
	const diff = end.getTime() - start.getTime();
	const newDiff = diff * Math.random();
	return new Date(start.getTime() + newDiff);
}
