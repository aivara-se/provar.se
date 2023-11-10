import type { Fetcher } from './index.js';

/**
 * MockFetcher is a mock implementation of the Fetcher interface.
 */
export class MockFetcher implements Fetcher {
	requests: Record<string, object[]>;
	responses: Record<string, object>;

	constructor() {
		this.requests = {};
		this.responses = {};
	}

	async fetch<TRes, TReq extends object = object>(
		method: string,
		path: string,
		body: TReq
	): Promise<TRes> {
		const key = this.createKey(method, path);
		if (this.responses[key]) {
			this.requests[key] = this.requests[key] || [];
			this.requests[key].push(body);
			return Promise.resolve(this.responses[key]) as Promise<TRes>;
		}
		throw new Error(`MockFetcher: No mock response found for ${key}`);
	}

	mockRequest(method: string, path: string, response: object) {
		const key = `${method.toUpperCase()} ${path}`;
		this.responses[key] = response;
	}

	getRequests(method: string, path: string): object[] {
		const key = this.createKey(method, path);
		return this.requests[key] || [];
	}

	private createKey(method: string, path: string) {
		return `${method.toUpperCase()} ${path}`;
	}
}
