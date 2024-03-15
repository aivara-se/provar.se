import {
	createAuthenticationEndpoints,
	createEmailAuthenticationEndpoints,
	createHealthCheckEndpoints,
	createOrganizationEndpoints
} from './api.code.js';
import { DefaultFetcher, Fetcher } from './fetcher.js';

/**
 * ClientOptions is the configuration object for the Client.
 */
export interface ClientOptions {
	token: string;
	fetcher?: Fetcher;
}

/**
 * ProvarClient is the main class for interacting with the API.
 */
export class ProvarClient {
	/**
	 * Grouped API endpoints
	 */
	EmailAuthenticationEndpoints;
	AuthenticationEndpoints;
	OrganizationEndpoints;
	HealthCheckEndpoints;

	/**
	 * The fetcher to use for sending requests.
	 */
	private fetcher: Fetcher;

	/**
	 * Creates a new Client instance.
	 */
	constructor(options: ClientOptions) {
		this.fetcher = options.fetcher || new DefaultFetcher(options.token);
		this.EmailAuthenticationEndpoints = createEmailAuthenticationEndpoints(this.fetcher);
		this.AuthenticationEndpoints = createAuthenticationEndpoints(this.fetcher);
		this.OrganizationEndpoints = createOrganizationEndpoints(this.fetcher);
		this.HealthCheckEndpoints = createHealthCheckEndpoints(this.fetcher);
	}
}
