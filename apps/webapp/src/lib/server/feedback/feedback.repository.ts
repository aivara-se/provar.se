import { getMongoClient } from '$lib/server/database';
import { ObjectId, type Collection } from 'mongodb';
import type { Feedback } from './feedback.types';

/**
 * The name of the MongoDB collection for feedback.
 */
const COLLECTION_NAME = 'feedbacks';

/**
 * Type for the MongoDB document for feedback in db.
 */
type FeedbackDocument = {
    _id: ObjectId;
    organization: string[];
};

/**
 * Convert a feedback document from the db to Feedback.
 */
function fromDocument(doc: FeedbackDocument): Feedback {
    return {
        id: doc._id.toHexString(),
        organization: doc.organization,
    };
}

/**
 * Get the MongoDB collection for feedback.
 */
async function getCollection(): Promise<Collection<FeedbackDocument>> {
    const mongo = await getMongoClient();
    return mongo.db().collection(COLLECTION_NAME);
}

/**
 * Find feedback by id.
 */
export async function findById(id: string): Promise<Feedback | null> {
    const coll = await getCollection();
    const doc = await coll.findOne({ _id: new ObjectId(id) });
    return doc ? fromDocument(doc) : null;
}

/**
 * Find feedback by organization
 */
export async function findByMember(organizationId: string): Promise<Feedback[]> {
    const coll = await getCollection();
    const docs = await coll.find({ organization: organizationId }).toArray();
    return docs.map(fromDocument);
}