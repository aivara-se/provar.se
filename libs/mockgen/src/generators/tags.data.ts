export const POSSIBLE_TAGS = {
	version: () => {
		const segment = (n: number) => Math.floor(Math.random() * n);
		return `${segment(2)}.${segment(5)}.${segment(5)}`;
	},
	feature: () => {
		const values = ['signup', 'onboarding', 'settings'];
		return values[Math.floor(Math.random() * values.length)];
	},
	package: () => {
		const values = ['essential', 'premium'];
		return values[Math.floor(Math.random() * values.length)];
	}
};
