import * as preslint from '@provar/eslint-config';

export default [
	...preslint.typescript(),
	// Add more configs here
	{
		ignores: ['src/api.type.ts', 'src/api.code.ts']
	}
];
