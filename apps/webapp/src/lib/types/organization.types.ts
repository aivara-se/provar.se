/**
 * The organization acts as a container for collected data.
 */
export interface Organization {
	id: string;
	createdAt: Date;
	name: string;
	description: string;
	members: string[];
}
