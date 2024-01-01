export const POSSIBLE_TAGS = {
	version: () => {
		const segment = (n: number) => Math.floor(Math.random() * n);
		return `${segment(10)}.${segment(10)}.${segment(10)}`;
	},
	feature: () => {
		const values = ['feature-x', 'feature-y', 'feature-z'];
		return values[Math.floor(Math.random() * values.length)];
	},
	package: () => {
		const values = ['free', 'basic', 'premium', 'enterprise'];
		return values[Math.floor(Math.random() * values.length)];
	}
};
