import { getMongoClient } from '$lib/server/database';
import type { Invitation } from '$lib/types';
import { ObjectId, type Collection } from 'mongodb';
import { createGravatarLink } from '../gravatar';
import { createInvitationLink } from './invitation.utils';

/**
 * The name of the MongoDB collection for invitations.
 */
const COLLECTION_NAME = 'invitations';

/**
 * Type for the MongoDB document for Invitations in db.
 */
interface InvitationDocument {
	_id: ObjectId;
	createdAt: Date;
	acceptedAt?: Date;
	key: string;
	name: string;
	email: string;
	organizationId: ObjectId;
}

/**
 * Convert a invitation document from the db to Invitation.
 */
function fromDocument(doc: InvitationDocument): Invitation {
	return {
		id: doc._id.toHexString(),
		createdAt: doc.createdAt,
		acceptedAt: doc.acceptedAt,
		key: doc.key,
		name: doc.name,
		link: createInvitationLink(doc.key),
		email: doc.email,
		image: createGravatarLink(doc.email),
		organizationId: doc.organizationId.toHexString()
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
	name: string;
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
		createdAt: new Date(),
		key: data.key,
		name: data.name,
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
 * Find an Invitation by id.
 */
export async function findById(id: string): Promise<Invitation | null> {
	const coll = await getCollection();
	const doc = await coll.findOne({ _id: new ObjectId(id) });
	return doc ? fromDocument(doc) : null;
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
