import { getMongoClient } from '$lib/server/database';
import { ObjectId, type Collection } from 'mongodb';
import type { Feature } from './feature.types';

/**
 * The name of the MongoDB collection for features.
 */
const COLLECTION_NAME = 'features';

/**
 * Type for the MongoDB document for Features in db.
 */
interface FeatureDocument {
    _id: ObjectId;
    name: string;
    key: string;
    organizationId: ObjectId;
};

/**
 * Convert a feature document from the db to Feature.
 */
function fromDocument(doc: FeatureDocument): Feature {
    return {
        id: doc._id.toHexString(),
        name: doc.name,
        key: doc.key,
        organizationId: doc.organizationId.toHexString(),
    };
}

/**
 * Get the MongoDB collection for features.
 */
async function getCollection(): Promise<Collection<FeatureDocument>> {
    const mongo = await getMongoClient();
    return mongo.db().collection(COLLECTION_NAME);
}

/**
 * Find a feature by id.
 */
export async function findById(id: string): Promise<Feature | null> {
    const coll = await getCollection();
    const doc = await coll.findOne({ _id: new ObjectId(id) });
    return doc ? fromDocument(doc) : null;
}

/**
 * Find a feature by key.
 */
export async function findByKey(key: string): Promise<Feature | null> {
    const coll = await getCollection();
    const doc = await coll.findOne({ key });
    return doc ? fromDocument(doc) : null;
}

/**
 * Find a feature by organisation
 */
export async function findByOrganisation(organizationId: string): Promise<Feature[]> {
    const coll = await getCollection();
    const doc = await coll.find({ organizationId: new ObjectId(organizationId) }).toArray();
    return doc.map(fromDocument);
}

