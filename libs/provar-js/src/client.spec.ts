import { beforeEach, describe, expect, it } from 'vitest';
import { ProvarClient } from './client.js';
import { MockFetcher } from './fetcher.mock.js';

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

	it('sends a text feedback with required fields', async () => {
		fetcher.mockRequest('POST', '/feedback', { success: true });
		await client.sendText('test feedback');
		const requests = fetcher.getRequests('POST', '/feedback');
		expect(requests.length).toBe(1);
		expect(requests[0]).toEqual({ type: 'text', data: { text: 'test feedback' } });
	});

	it('sends a cNPS feedback with required fields', async () => {
		fetcher.mockRequest('POST', '/feedback', { success: true });
		await client.sendCNPS(0.8);
		const requests = fetcher.getRequests('POST', '/feedback');
		expect(requests.length).toBe(1);
		expect(requests[0]).toEqual({ type: 'cnps', data: { cnps: 0.8 } });
	});

	it('sends a CSAT feedback with required fields', async () => {
		fetcher.mockRequest('POST', '/feedback', { success: true });
		await client.sendCSAT(0.8);
		const requests = fetcher.getRequests('POST', '/feedback');
		expect(requests.length).toBe(1);
		expect(requests[0]).toEqual({ type: 'csat', data: { csat: 0.8 } });
	});

	it('sends a text feedback with all optional fields', async () => {
		fetcher.mockRequest('POST', '/feedback', { success: true });
		await client.sendText('test feedback', 'project1', { tag1: 'value1', tag2: 'value2' });
		const requests = fetcher.getRequests('POST', '/feedback');
		expect(requests.length).toBe(1);
		expect(requests[0]).toEqual({
			type: 'text',
			data: { text: 'test feedback' },
			projectId: 'project1',
			tags: { tag1: 'value1', tag2: 'value2' }
		});
	});
});
