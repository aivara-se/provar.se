import { defineConfig } from 'vitest/config';

export default defineConfig({
	test: {
		name: 'mockgen',
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
