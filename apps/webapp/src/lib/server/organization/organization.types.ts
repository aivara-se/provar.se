/**
 * The organization acts as a container for collected data.
 */
export interface Organization {
	id: string;
	slug: string;
	name: string;
	prod: boolean;
	members: string[];
}
