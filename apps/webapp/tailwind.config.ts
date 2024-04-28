import daisyui from 'daisyui';
import typography from '@tailwindcss/typography';
import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {}
	},
	plugins: [typography, daisyui],
	daisyui: {
		themes: ['dark'],
		logs: false
	}
} satisfies Config;
