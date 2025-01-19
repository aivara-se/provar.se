import * as preslint from '@provar/eslint-config';

export default [
	...preslint.typescript(),

	// Ignore auto-generated files and build files from linting
	{ ignores: ['src/api.type.ts', 'src/api.code.ts', 'build/*'] }
];
