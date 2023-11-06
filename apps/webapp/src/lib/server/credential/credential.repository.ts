import { getMongoClient } from '$lib/server/database';
import { ObjectId, type Collection } from 'mongodb';
import type { Credential } from './credential.types';

/**
 * The name of the MongoDB collection for credentials.
 */
const COLLECTION_NAME = 'credentials';

/**
 * Type for the MongoDB document for credentials in db.
 */
type CredentialDocument = {
	_id: ObjectId;
	name: string;
	organizationId: ObjectId;
	clientId: string;
	clientSecret: string;
};

/**
 * Convert an credential document from the db to Credential.
 */
function fromDocument(doc: CredentialDocument): Credential {
	return {
		id: doc._id.toHexString(),
		name: doc.name,
		organizationId: doc.organizationId.toHexString(),
		clientId: doc.clientId,
		clientSecret: doc.clientSecret
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
	clientId: string;
	clientSecret: string;
}

/**
 * Create a new credential in the database with the given information.
 */
export async function create(data: CreateCredentialData): Promise<void> {
	const coll = await getCollection();
	await coll.insertOne({
		_id: new ObjectId(),
		name: data.name,
		organizationId: new ObjectId(data.organizationId),
		clientId: data.clientId,
		clientSecret: data.clientSecret
	});
}

/**
 * Delete a credential by id.
 */
export async function deleteById(id: string): Promise<void> {
	const coll = await getCollection();
	await coll.deleteOne({ _id: new ObjectId(id) });
}

/**
 * Find credentials by organization.
 */
export async function findByOrganization(organizationId: string): Promise<Credential[]> {
	const coll = await getCollection();
	const docs = await coll.find({ organizationId: new ObjectId(organizationId) }).toArray();
	return docs.map(fromDocument);
}
