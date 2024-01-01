export function generateRandomDate(start: Date, end: Date): Date {
	const diff = end.getTime() - start.getTime();
	const newDiff = diff * Math.random();
	return new Date(start.getTime() + newDiff);
}
