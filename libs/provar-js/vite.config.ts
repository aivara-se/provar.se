import { defineConfig } from 'vitest/config';

export default defineConfig({
	test: {
		name: 'provar-js',
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
