import fetch from 'isomorphic-fetch';
import { config } from './config.js';

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
		if (fetchResponse.status === 204) {
			return null as TRes;
		}
		return fetchResponse.json();
	}
}
