import { getMongoClient } from '$lib/server/database';
import { ObjectId, type Collection } from 'mongodb';
import type { Feature } from './features.types';

/**
 * The name of the MongoDB collection for features.
 */
const COLLECTION_NAME = 'features';

/**
 * Type for the MongoDB document for Features in db.
 */
type FeatureDocument = {
    _id: ObjectId;
    name: string;
};

/**
 * Convert a feature document from the db to Feature.
 */
function fromDocument(doc: FeatureDocument): Feature {
    return {
        id: doc._id.toHexString(),
        name: doc.name,
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
 * Find a feature by name.
 */
export async function findByName(name: string): Promise<Feature | null> {
    const coll = await getCollection();
    const doc = await coll.findOne({ name });
    return doc ? fromDocument(doc) : null;
}
