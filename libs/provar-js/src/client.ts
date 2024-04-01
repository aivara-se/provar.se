import {
	createAuthenticationEndpoints,
	createCredentialEndpoints,
	createEmailAuthenticationEndpoints,
	createFeedbackEndpoints,
	createHealthCheckEndpoints,
	createInvitationEndpoints,
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
	Authentication;
	Credential;
	EmailAuthentication;
	Feedback;
	HealthCheck;
	Invitation;
	Organization;

	/**
	 * The fetcher to use for sending requests.
	 */
	private fetcher: Fetcher;

	/**
	 * Creates a new Client instance.
	 */
	constructor(options: ClientOptions) {
		this.fetcher = options.fetcher || new DefaultFetcher(options.token);
		this.Authentication = createAuthenticationEndpoints(this.fetcher);
		this.Credential = createCredentialEndpoints(this.fetcher);
		this.EmailAuthentication = createEmailAuthenticationEndpoints(this.fetcher);
		this.Feedback = createFeedbackEndpoints(this.fetcher);
		this.HealthCheck = createHealthCheckEndpoints(this.fetcher);
		this.Invitation = createInvitationEndpoints(this.fetcher);
		this.Organization = createOrganizationEndpoints(this.fetcher);
	}
}
