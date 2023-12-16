import { getMongoClient } from '$lib/server/database';
import { ObjectId, type Collection } from 'mongodb';
import type { Organization } from '$lib/types';

/**
 * The name of the MongoDB collection for organizations.
 */
const COLLECTION_NAME = 'organizations';

/**
 * Type for the MongoDB document for organizations in db.
 */
type OrganizationDocument = {
	_id: ObjectId;
	name: string;
	description?: string;
	members: ObjectId[];
};

/**
 * Convert an organization document from the db to Organization.
 */
function fromDocument(doc: OrganizationDocument): Organization {
	return {
		id: doc._id.toHexString(),
		name: doc.name,
		description: doc.description ?? '',
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
	name: string;
	description?: string;
}

/**
 * Create a new organization in the database with the given information.
 */
export async function create(userId: string, data: CreateOrganizationData): Promise<string> {
	const coll = await getCollection();
	const id = new ObjectId();
	await coll.insertOne({
		_id: id,
		name: data.name,
		description: '',
		members: [new ObjectId(userId)]
	});
	return id.toHexString();
}

/**
 * The data required to update an organization.
 */
export interface UpdateOrganizationData {
	name?: string;
	description?: string;
}

/**
 * Update organization in the database with the given information.
 */
export async function update(organizationId: string, data: UpdateOrganizationData): Promise<void> {
	const coll = await getCollection();
	await coll.updateOne({ _id: new ObjectId(organizationId) }, { $set: data });
}

/**
 * Removes an organization in the organizations collection.
 */
export async function remove(organizationId: string): Promise<void> {
	const coll = await getCollection();
	await coll.deleteOne({ _id: new ObjectId(organizationId) });
}

/**
 * Removes a member from an organization.
 */
export async function removeMember(organizationId: string, userId: string): Promise<void> {
	const coll = await getCollection();
	await coll.updateOne(
		{ _id: new ObjectId(organizationId) },
		{ $pull: { members: new ObjectId(userId) } }
	);
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
	const docs = await coll.find({ members: new ObjectId(userId) }, { sort: { name: 1 } }).toArray();
	return docs.map(fromDocument);
}
