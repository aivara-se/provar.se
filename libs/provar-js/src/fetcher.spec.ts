import { beforeEach, describe, it, expect, vi } from 'vitest';
import { DefaultFetcher } from './fetcher';

describe('DefaultFetcher', () => {
	let fetcher: DefaultFetcher;

	beforeEach(() => {
		fetcher = new DefaultFetcher('token');
	});

	it('should set the correct headers', () => {
		expect(fetcher['headers']).toEqual({
			'Content-Type': 'application/json',
			Authorization: 'Bearer token'
		});
	});

	it('should make a GET request', async () => {
		const mockResponse = { data: 'response' };
		const mockFetch = vi.fn().mockResolvedValue({
			status: 200,
			json: vi.fn().mockResolvedValue(mockResponse)
		});
		fetcher['_fetch'] = mockFetch;

		const result = await fetcher.fetch('GET', '/path');

		expect(mockFetch).toHaveBeenCalledWith('https://api.provar.se/path', {
			method: 'GET',
			headers: fetcher['headers']
		});
		expect(result).toEqual(mockResponse);
	});

	it('should make a POST request', async () => {
		const mockResponse = { data: 'response' };
		const mockFetch = vi.fn().mockResolvedValue({
			status: 200,
			json: vi.fn().mockResolvedValue(mockResponse)
		});
		fetcher['_fetch'] = mockFetch;

		const result = await fetcher.fetch('POST', '/path', { body: 'data' });

		expect(mockFetch).toHaveBeenCalledWith('https://api.provar.se/path', {
			method: 'POST',
			headers: fetcher['headers'],
			body: JSON.stringify({ body: 'data' })
		});
		expect(result).toEqual(mockResponse);
	});

	it('should throw an error for non-success status code', async () => {
		const mockFetch = vi.fn().mockResolvedValue({
			status: 404
		});
		fetcher['_fetch'] = mockFetch;

		await expect(fetcher.fetch('GET', '/path')).rejects.toThrowError(
			'Provar: request failed with status 404'
		);
	});

	it('should return null for status code 204', async () => {
		const mockFetch = vi.fn().mockResolvedValue({
			status: 204
		});
		fetcher['_fetch'] = mockFetch;

		const result = await fetcher.fetch('GET', '/path');

		expect(result).toBeNull();
	});
});
