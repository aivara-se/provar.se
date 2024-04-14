import adapter from '@sveltejs/adapter-cloudflare';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: vitePreprocess(),

	kit: {
		// Use default settings for cloudflare adapter
		//  * https://kit.svelte.dev/docs/single-page-apps
		//  * https://kit.svelte.dev/docs/adapter-cloudflare
		adapter: adapter()
	}
};

export default config;
