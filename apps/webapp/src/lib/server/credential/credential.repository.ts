import { getMongoClient } from '$lib/server/database';
import { ObjectId, type Collection } from 'mongodb';
import type { Credential } from '$lib/types';

/**
 * The name of the MongoDB collection for credentials.
 */
const COLLECTION_NAME = 'credentials';

/**
 * Type for the MongoDB document for credentials in db.
 */
type CredentialDocument = {
	_id: ObjectId;
	createdAt: Date;
	lastUsedAt?: Date;
	name: string;
	organizationId: ObjectId;
	key: string;
};

/**
 * Convert an credential document from the db to Credential.
 */
function fromDocument(doc: CredentialDocument): Credential {
	return {
		id: doc._id.toHexString(),
		createdAt: doc.createdAt,
		lastUsedAt: doc.lastUsedAt,
		name: doc.name,
		organizationId: doc.organizationId.toHexString(),
		key: doc.key
	};
}

/**
 * Get the MongoDB collection for credentials.
 */
async function getCollection(): Promise<Collection<CredentialDocument>> {
	const mongo = await getMongoClient();
	return mongo.db().collection(COLLECTION_NAME);
}

/**
 * The data required to create a new credential.
 */
export interface CreateCredentialData {
	name: string;
	organizationId: string;
	key: string;
}

/**
 * Create a new credential in the database with the given information.
 */
export async function create(data: CreateCredentialData): Promise<string> {
	const coll = await getCollection();
	const id = new ObjectId();
	await coll.insertOne({
		_id: id,
		createdAt: new Date(),
		name: data.name,
		organizationId: new ObjectId(data.organizationId),
		key: data.key
	});
	return id.toHexString();
}

/**
 * Revoke a credential by id.
 */
export async function remove(organizationId: string, id: string): Promise<void> {
	const coll = await getCollection();
	await coll.deleteOne({
		_id: new ObjectId(id),
		organizationId: new ObjectId(organizationId)
	});
}

/**
 * Revoke all credentials that belong to an organization
 */
export async function removeAll(organizationId: string): Promise<void> {
	const coll = await getCollection();
	await coll.deleteMany({ organizationId: new ObjectId(organizationId) });
}

/**
 * Find a credential by id.
 */
export async function findById(organizationId: string, id: string): Promise<Credential | null> {
	const coll = await getCollection();
	const doc = await coll.findOne({
		_id: new ObjectId(id),
		organizationId: new ObjectId(organizationId)
	});
	return doc ? fromDocument(doc) : null;
}

/**
 * Find credentials by organization.
 */
export async function findByOrganization(organizationId: string): Promise<Credential[]> {
	const coll = await getCollection();
	const docs = await coll
		.find({ organizationId: new ObjectId(organizationId) }, { sort: { name: 1 } })
		.toArray();
	return docs.map(fromDocument);
}
