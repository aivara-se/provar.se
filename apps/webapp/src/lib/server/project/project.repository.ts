import { getMongoClient } from '$lib/server/database';
import type { Project } from '$lib/types';
import { ObjectId, type Collection } from 'mongodb';

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
	description?: string;
	organizationId: ObjectId;
	collectionGoal?: number;
}

/**
 * Convert a project document from the db to Project.
 */
function fromDocument(doc: ProjectDocument): Project {
	return {
		id: doc._id.toHexString(),
		name: doc.name,
		description: doc.description ?? '',
		organizationId: doc.organizationId.toHexString(),
		collectionGoal: doc.collectionGoal ?? null
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
 * Update project information
 */
export interface UpdateProjectData {
	name: string;
}

/**
 * Update project information in the database
 */
export async function update(
	organizationId: string,
	projectId: string,
	data: UpdateProjectData
): Promise<void> {
	const coll = await getCollection();
	await coll.updateOne(
		{ _id: new ObjectId(projectId), organizationId: new ObjectId(organizationId) },
		{ $set: { name: data.name } }
	);
}

/**
 * Removes a project in the projects collection.
 */
export async function remove(organizationId: string, projectId: string): Promise<void> {
	const coll = await getCollection();
	await coll.deleteOne({
		_id: new ObjectId(projectId),
		organizationId: new ObjectId(organizationId)
	});
}

/**
 * Remove all projects that belong to an organization.
 */
export async function removeAll(organizationId: string): Promise<void> {
	const coll = await getCollection();
	await coll.deleteMany({ organizationId: new ObjectId(organizationId) });
}

/**
 * Find a Project by id.
 */
export async function findById(organizationId: string, projectId: string): Promise<Project | null> {
	const coll = await getCollection();
	const doc = await coll.findOne({
		_id: new ObjectId(projectId),
		organizationId: new ObjectId(organizationId)
	});
	return doc ? fromDocument(doc) : null;
}

/**
 * Find project by organization.
 */
export async function findByOrganization(organizationId: string): Promise<Project[]> {
	const coll = await getCollection();
	const docs = await coll
		.find({ organizationId: new ObjectId(organizationId) }, { sort: { name: -1 } })
		.toArray();
	return docs.map(fromDocument);
}
