import { getMongoClient } from '$lib/server/database';
import type { Invitation } from '$lib/types';
import { ObjectId, type Collection } from 'mongodb';

/**
 * The name of the MongoDB collection for invitations.
 */
const COLLECTION_NAME = 'invitations';

/**
 * Type for the MongoDB document for Invitations in db.
 */
interface InvitationDocument {
	_id: ObjectId;
	key: string;
	email: string;
	organizationId: ObjectId;
}

/**
 * Convert a invitation document from the db to Invitation.
 */
function fromDocument(doc: InvitationDocument): Invitation {
	return {
		id: doc._id.toHexString(),
		key: doc.key,
		email: doc.email,
		organizationId: doc.organizationId.toHexString(),
		createdAt: doc._id.getTimestamp()
	};
}

/**
 * Get the MongoDB collection for invitations.
 */
async function getCollection(): Promise<Collection<InvitationDocument>> {
	const mongo = await getMongoClient();
	return mongo.db().collection(COLLECTION_NAME);
}

/**
 * The data required to create a new invitation.
 */
export interface CreateInvitationData {
	key: string;
	email: string;
	organizationId: string;
}

/**
 * Create a new invitation in the database with the given information.
 */
export async function create(data: CreateInvitationData): Promise<string> {
	const coll = await getCollection();
	const id = new ObjectId();
	await coll.insertOne({
		_id: id,
		key: data.key,
		email: data.email,
		organizationId: new ObjectId(data.organizationId)
	});
	return id.toHexString();
}

/**
 * Removes a invitation by id.
 */
export async function remove(organizationId: string, invitationId: string): Promise<void> {
	const coll = await getCollection();
	await coll.deleteOne({
		_id: new ObjectId(invitationId),
		organizationId: new ObjectId(organizationId)
	});
}

/**
 * Find an Invitation by key.
 */
export async function findByKey(invitationKey: string): Promise<Invitation | null> {
	const coll = await getCollection();
	const doc = await coll.findOne({ key: invitationKey });
	return doc ? fromDocument(doc) : null;
}

/**
 * Find invitations by organization.
 */
export async function findByOrganization(organizationId: string): Promise<Invitation[]> {
	const coll = await getCollection();
	const docs = await coll.find({ organizationId: new ObjectId(organizationId) }).toArray();
	return docs.map(fromDocument);
}
