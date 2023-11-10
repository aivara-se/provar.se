import { beforeEach, describe, expect, it } from 'vitest';
import { ProvarClient } from './index.js';
import { MockFetcher } from './index.mock.js';

/**
 * ProvarClient is the main class for interacting with the API.
 */
describe('ProvarClient', () => {
	let client: ProvarClient;
	let fetcher: MockFetcher;

	beforeEach(() => {
		fetcher = new MockFetcher();
		client = new ProvarClient({
			apiKey: 'test-api-key',
			fetcher: fetcher
		});
	});

	it('sends feedback successfully', async () => {
		fetcher.mockRequest('POST', '/feedback', { success: true });
		await client.sendFeedback({ text: 'test feedback' });
		const requests = fetcher.getRequests('POST', '/feedback');
		expect(requests.length).toBe(1);
		expect(requests[0]).toEqual({ text: 'test feedback' });
	});
});
