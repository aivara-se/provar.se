// @ts-check

import eslint from '@eslint/js';
import tseslint from 'typescript-eslint';

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
