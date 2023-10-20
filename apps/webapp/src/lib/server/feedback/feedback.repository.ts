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
interface FeedbackDocument {
    _id: ObjectId;
    organizationId: ObjectId;
    projectId: ObjectId;
};

/**
 * Convert a feedback document from the db to Feedback.
 */
function fromDocument(doc: FeedbackDocument): Feedback {
    return {
        id: doc._id.toHexString(),
        organizationId: doc.organizationId.toHexString(),
        projectId: doc.projectId.toHexString(),
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
 * Find feedback by organization.
 */
export async function findByOrganization(organizationId: string): Promise<Feedback[]> {
    const coll = await getCollection();
    const docs = await coll.find({ organizationId: new ObjectId(organizationId) }).toArray();
    return docs.map(fromDocument);
}

/**
 * Find feedback by project.
 */
export async function findByProject(projectId: string, organizationId: string): Promise<Feedback[]> {
    const coll = await getCollection();
    const docs = await coll.find({ projectId: new ObjectId(projectId), organizationId: new ObjectId(organizationId) }).toArray();
    return docs.map(fromDocument);
}