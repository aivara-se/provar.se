import { getMongoClient } from '$lib/server/database';
import { ObjectId, type Collection } from 'mongodb';
import type { Organization } from './organization.types';

/**
 * The name of the MongoDB collection for organizations.
 */
const COLLECTION_NAME = 'organizations';

/**
 * Type for the MongoDB document for organizations in db.
 */
type OrganizationDocument = {
	_id: ObjectId;
	slug: string;
	name: string;
	prod: boolean;
	members: ObjectId[];
};

/**
 * Convert an organization document from the db to Organization.
 */
function fromDocument(doc: OrganizationDocument): Organization {
	return {
		id: doc._id.toHexString(),
		slug: doc.slug,
		name: doc.name,
		prod: !!doc.prod,
		members: doc.members.map((id) => id.toHexString())
	};
}

/**
 * Get the MongoDB collection for organizations.
 */
async function getCollection(): Promise<Collection<OrganizationDocument>> {
	const mongo = await getMongoClient();
	return mongo.db().collection(COLLECTION_NAME);
}

/**
 * The data required to create a new organization.
 */
export interface CreateOrganizationData {
	slug: string;
	name: string;
	prod: boolean;
}

/**
 * Create a new organization in the database with the given information.
 */
export async function create(userId: string, data: CreateOrganizationData): Promise<void> {
	const coll = await getCollection();
	await coll.insertOne({
		_id: new ObjectId(),
		slug: data.slug,
		name: data.name,
		prod: data.prod,
		members: [new ObjectId(userId)]
	});
}

/**
 * Find an organization by id.
 */
export async function findById(id: string): Promise<Organization | null> {
	const coll = await getCollection();
	const doc = await coll.findOne({ _id: new ObjectId(id) });
	return doc ? fromDocument(doc) : null;
}

/**
 * Get a list of all organizations that the given user is a member of.
 */
export async function findByMember(userId: string): Promise<Organization[]> {
	const coll = await getCollection();
	const docs = await coll.find({ members: new ObjectId(userId) }).toArray();
	return docs.map(fromDocument);
}
