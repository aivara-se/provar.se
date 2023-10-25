import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';

export default defineConfig({
	plugins: [sveltekit()],
	test: {
		name: '@aivara-se/provar-webapp',
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
