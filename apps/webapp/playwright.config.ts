import type { PlaywrightTestConfig } from '@playwright/test';

const config: PlaywrightTestConfig = {
	webServer: {
		command: 'yarn preview',
		port: 4173
	},
	testDir: 'tests',
	testMatch: /.+\.integration\.spec\.ts/
};

export default config;
