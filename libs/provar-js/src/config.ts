/**
 * config is the configuration for the library.
 */
export const config = {
	baseUrl: 'https://api.provar.se'
};

/**
 * _overrideConfig is a helper function for overriding the default config values.
 * This is used for testing. The library should not be used with this function.
 */
export function _overrideConfig(overrides: Partial<typeof config>): void {
	Object.assign(config, overrides);
}
