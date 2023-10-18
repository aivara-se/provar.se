import { getMongoClient } from '$lib/server/database';
import { ObjectId, type Collection } from 'mongodb';
import type { User } from './user.types';

/**
 * The name of the MongoDB collection for users.
 */
const COLLECTION_NAME = 'users';

/**
 * Type for the MongoDB document for users in db.
 */
type UserDocument = {
	_id: ObjectId;
	name: string;
	email: string;
	image: string;
	emailVerified: Date | null;
};

/**
 * Convert a user document from the db to User.
 */
function fromDocument(doc: UserDocument): User {
	return {
		id: doc._id.toHexString(),
		name: doc.name,
		email: doc.email,
		image: doc.image
	};
}

/**
 * Get the MongoDB collection for users.
 */
async function getCollection(): Promise<Collection<UserDocument>> {
	const mongo = await getMongoClient();
	return mongo.db().collection(COLLECTION_NAME);
}

/**
 * Find a user by id.
 */
export async function findById(id: string): Promise<User | null> {
	const coll = await getCollection();
	const doc = await coll.findOne({ _id: new ObjectId(id) });
	return doc ? fromDocument(doc) : null;
}

/**
 * Find a user by email.
 */
export async function findByEmail(email: string): Promise<User | null> {
	const coll = await getCollection();
	const doc = await coll.findOne({ email });
	return doc ? fromDocument(doc) : null;
}
