import { DefaultFetcher, Fetcher } from './fetcher.js';
import {
	CNPSFeedback,
	CSATFeedback,
	Feedback,
	FeedbackTags,
	FeedbackType,
	TextFeedback
} from './types.js';

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
	 * Sends a text feedback message to the Provar API.
	 */
	async sendText(text: string, projectId?: string, tags?: FeedbackTags): Promise<void> {
		const feedback: TextFeedback = { type: FeedbackType.Text, data: { text } };
		await this.send(feedback, projectId, tags);
	}

	/**
	 * Sends a cNPS feedback message to the Provar API.
	 */
	async sendCNPS(cnps: number, projectId?: string, tags?: FeedbackTags): Promise<void> {
		const feedback: CNPSFeedback = { type: FeedbackType.CNPS, data: { cnps } };
		await this.send(feedback, projectId, tags);
	}

	/**
	 * Sends a CSAT feedback message to the Provar API.
	 */
	async sendCSAT(csat: number, projectId?: string, tags?: FeedbackTags): Promise<void> {
		const feedback: CSATFeedback = { type: FeedbackType.CSAT, data: { csat } };
		await this.send(feedback, projectId, tags);
	}

	/**
	 * Sends a feedback message to the Provar API.
	 */
	private async send(feedback: Feedback, projectId?: string, tags?: FeedbackTags): Promise<void> {
		if (projectId) {
			feedback.projectId = projectId;
		}
		if (tags) {
			feedback.tags = tags;
		}
		await this.fetcher.fetch('post', '/feedback', feedback);
	}
}
