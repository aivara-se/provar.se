import fetch from 'isomorphic-fetch';
import type { Feedback } from './types.js';

/**
 * config is the configuration for the library.
 */
const config = {
	baseUrl: 'https://api.provar.se'
};

/**
 * Fetcher is the interface for classes sending requests to the Provar API.
 */
export interface Fetcher {
	fetch<TRes, TReq extends object>(method: string, path: string, body: TReq): Promise<TRes>;
}

/**
 * DefaultFetcher is the default fetcher to use for sending requests.
 */
export class DefaultFetcher implements Fetcher {
	/**
	 * HTTP headers sent with each request.
	 */
	private headers: Record<string, string>;

	/**
	 * Creates a new ProvarClient instance.
	 */
	constructor(apiKey: string) {
		this.headers = {
			Accept: 'application/json',
			'Content-Type': 'application/json',
			Authorization: `Bearer ${apiKey}`
		};
	}

	/**
	 * Sends a request to the Provar API with the auth header and JSON headers.
	 * Throws an error if the request fails or returns a non-success status code.
	 */
	async fetch<TRes, TReq extends object = object>(
		method: string,
		path: string,
		body: TReq
	): Promise<TRes> {
		const fetchResponse = await fetch(`${config.baseUrl}${path}`, {
			method,
			headers: this.headers,
			body: JSON.stringify(body)
		});
		if (fetchResponse.status >= 400) {
			throw new Error(`Provar: request failed with status ${fetchResponse.status}`);
		}
		return fetchResponse.json();
	}
}

/**
 * ClientOptions is the configuration object for the Client.
 */
export interface ClientOptions {
	apiKey: string;
	fetcher?: Fetcher;
}

/**
 * ProvarClient is the main class for interacting with the API.
 */
export class ProvarClient {
	/**
	 * The fetcher to use for sending requests.
	 */
	private fetcher: Fetcher;

	/**
	 * Creates a new Client instance.
	 */
	constructor(options: ClientOptions) {
		this.fetcher = options.fetcher || new DefaultFetcher(options.apiKey);
	}

	/**
	 * Sends a feedback message to the Provar API.
	 */
	async sendFeedback(feedback: Feedback): Promise<void> {
		await this.fetcher.fetch('post', '/feedback', feedback);
	}
}

/**
 * _overrideConfig is a helper function for overriding the default config values.
 * This is used for testing. The library should not be used with this function.
 */
export function _overrideConfig(overrides: Partial<typeof config>): void {
	Object.assign(config, overrides);
}
