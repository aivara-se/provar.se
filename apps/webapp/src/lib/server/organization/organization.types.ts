/**
 * Is the organization on production or staging environment?
 */
export type Environment = 'staging' | 'production';

/**
 * The organization acts as a container for collected data.
 */
export interface Organization {
	id: string;
	slug: string;
	name: string;
	environment: Environment;
	members: string[];
}
