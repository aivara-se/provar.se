import { getMongoClient } from '$lib/server/database';
import { ObjectId, type Collection } from 'mongodb';
import type { Project } from './project.types';

/**
 * The name of the MongoDB collection for projects.
 */
const COLLECTION_NAME = 'projects';

/**
 * Type for the MongoDB document for Projects in db.
 */
interface ProjectDocument {
	_id: ObjectId;
	name: string;
	organizationId: ObjectId;
}

/**
 * Convert a project document from the db to Project.
 */
function fromDocument(doc: ProjectDocument): Project {
	return {
		id: doc._id.toHexString(),
		name: doc.name,
		organizationId: doc.organizationId.toHexString()
	};
}

/**
 * Get the MongoDB collection for projects.
 */
async function getCollection(): Promise<Collection<ProjectDocument>> {
	const mongo = await getMongoClient();
	return mongo.db().collection(COLLECTION_NAME);
}

/**
 * The data required to create a new project.
 */
export interface CreateProjectData {
	name: string;
	organizationId: string;
}

/**
 * Create a new project in the database with the given information.
 */
export async function create(data: CreateProjectData): Promise<string> {
	const coll = await getCollection();
	const id = new ObjectId();
	await coll.insertOne({
		_id: id,
		name: data.name,
		organizationId: new ObjectId(data.organizationId)
	});
	return id.toHexString();
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
 * Find project by organization.
 */
export async function findByOrganization(organizationId: string): Promise<Project[]> {
	const coll = await getCollection();
	const docs = await coll.find({ organizationId: new ObjectId(organizationId) }).toArray();
	return docs.map(fromDocument);
}
