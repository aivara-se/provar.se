import merge from 'lodash.merge';
import { DefaultFetcher, Fetcher } from './fetcher.js';
import { Feedback, FeedbackMeta, FeedbackTags, FeedbackType, FeedbackUser } from './types.js';

/**
 * ClientOptions is the configuration object for the Client.
 */
export interface ClientOptions {
	apiKey: string;
	fetcher?: Fetcher;
}

/**
 * FeedbackMetadata is additinonal options for sending feedback.
 */
export type FeedbackMetadata = Omit<Partial<Feedback>, 'type' | 'data'>;

/**
 * ProvarClient is the main class for interacting with the API.
 */
export class ProvarClient {
	/**
	 * The fetcher to use for sending requests.
	 */
	private fetcher: Fetcher;

	/**
	 * The fetcher to use for sending requests.
	 */
	private globals: FeedbackMetadata = {};

	/**
	 * Creates a new Client instance.
	 */
	constructor(options: ClientOptions) {
		this.fetcher = options.fetcher || new DefaultFetcher(options.apiKey);
	}

	/**
	 * Sets the project ID field sent with all feedback messages.
	 */
	setProject(projectId: string): void {
		this.globals.projectId = projectId;
	}

	/**
	 * Sets the user field sent with all feedback messages.
	 */
	setUser(user: FeedbackUser): void {
		this.globals.user = user;
	}

	/**
	 * Sets the meta field sent with all feedback messages.
	 */
	setMeta(meta: FeedbackMeta): void {
		this.globals.meta = { ...this.globals.meta, ...meta };
	}

	/**
	 * Sets the tags field sent with all feedback messages.
	 */
	setTags(tags: FeedbackTags): void {
		this.globals.tags = { ...this.globals.tags, ...tags };
	}

	/**
	 * Sends a text feedback message to the Provar API.
	 */
	async sendText(text: string, data: FeedbackMetadata = {}): Promise<void> {
		await this.send({ ...data, type: FeedbackType.Text, data: { text } });
	}

	/**
	 * Sends a cNPS feedback message to the Provar API.
	 */
	async sendCNPS(cnps: number, data: FeedbackMetadata = {}): Promise<void> {
		await this.send({ ...data, type: FeedbackType.CNPS, data: { cnps } });
	}

	/**
	 * Sends a CSAT feedback message to the Provar API.
	 */
	async sendCSAT(csat: number, data: FeedbackMetadata = {}): Promise<void> {
		await this.send({ ...data, type: FeedbackType.CSAT, data: { csat } });
	}

	/**
	 * Sends a feedback message to the Provar API.
	 */
	private async send(feedback: Feedback): Promise<void> {
		const payload = merge({}, this.globals, feedback);
		await this.fetcher.fetch('post', '/feedback', payload);
	}
}
