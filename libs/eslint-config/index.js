// @ts-check

import eslint from '@eslint/js';
import tseslint from 'typescript-eslint';
import sveslint from 'eslint-plugin-svelte';

/**
 * typescript returns eslint configs for typescript libraries.
 */
export function typescript() {
	return [
		// typescript-eslint configs
		...tseslint.config(
			eslint.configs.recommended,
			tseslint.configs.strict,
			tseslint.configs.stylistic
		)
	];
}

/**
 * sveltekit returns eslint configs for SvelteKit applications.
 */
export function sveltekit() {
	return [
		...typescript(),
		...sveslint.configs['flat/recommended'],
		...sveslint.configs['flat/prettier']
	];
}
