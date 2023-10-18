import { getMongoClient } from '$lib/server/database';
import { ObjectId, type Collection } from 'mongodb';
import type { Project } from './project.types';

/**
 * The name of the MongoDB collection for projects.
 */
const COLLECTION_NAME = 'Project';

/**
 * Type for the MongoDB document for Projects in db.
 */
type projectDocument = {
    _id: ObjectId;
    name: string;
};

/**
 * Convert a project document from the db to Project.
 */
function fromDocument(doc: projectDocument): Project {
    return {
        id: doc._id.toHexString(),
        name: doc.name,
    };
}

/**
 * Get the MongoDB collection for projects.
 */
async function getCollection(): Promise<Collection<projectDocument>> {
    const mongo = await getMongoClient();
    return mongo.db().collection(COLLECTION_NAME);
}

/**
 * Find a Project by id.
 */
export async function findById(id: string): Promise<Project | null> {
    const coll = await getCollection();
    const doc = await coll.findOne({ _id: new ObjectId(id) });
    return doc ? fromDocument(doc) : null;
}

/**
 * Find a project by name.
 */
export async function findByName(name: string): Promise<Project | null> {
    const coll = await getCollection();
    const doc = await coll.findOne({ name });
    return doc ? fromDocument(doc) : null;
}

