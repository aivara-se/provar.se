import { getMongoClient } from '$lib/server/database';
import type {
	FeedbackType,
	Feedback,
	TextFeedbackData,
	CNPSFeedbackData,
	CSATFeedbackData
} from '$lib/types';
import { ObjectId, type Collection } from 'mongodb';

/**
 * The name of the MongoDB collection for feedback.
 */
const COLLECTION_NAME = 'feedbacks';

/**
 * Base type for the MongoDB document for feedback in db. This should not
 * be used. Use derived types instead.
 */
interface BaseFeedbackDocument {
	_id: ObjectId;
	organizationId: ObjectId;
	projectId?: ObjectId;
	type: FeedbackType;
	tags?: Record<string, string>;
	data: unknown;
}

/**
 * Type for the variant of the document that stores text feedback.
 */
interface TextFeedbackDocument extends BaseFeedbackDocument {
	type: FeedbackType.Text;
	data: TextFeedbackData;
}

/**
 * Type for the variant of the document that stores cNPS feedback.
 */
interface CNPSFeedbackDocument extends BaseFeedbackDocument {
	type: FeedbackType.CNPS;
	data: CNPSFeedbackData;
}

/**
 * Type for the variant of the document that stores CSAT feedback.
 */
interface CSATFeedbackDocument extends BaseFeedbackDocument {
	type: FeedbackType.CSAT;
	data: CSATFeedbackData;
}

/**
 * Type for the MongoDB document for feedback in db.
 */
type FeedbackDocument = TextFeedbackDocument | CNPSFeedbackDocument | CSATFeedbackDocument;

/**
 * Convert a feedback document from the db to Feedback.
 */
function fromDocument(doc: FeedbackDocument): Feedback {
	return {
		id: doc._id.toHexString(),
		organizationId: doc.organizationId.toHexString(),
		projectId: doc.projectId?.toHexString(),
		createdAt: doc._id.getTimestamp(),
		type: doc.type,
		// TODO: Fix this any without writing unnecessary javascript code.
		// eslint-disable-next-line @typescript-eslint/no-explicit-any
		data: doc.data as any,
		tags: doc.tags || {}
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
export async function findById(organizationId: string, id: string): Promise<Feedback | null> {
	const coll = await getCollection();
	const doc = await coll.findOne({
		organizationId: new ObjectId(organizationId),
		_id: new ObjectId(id)
	});
	return doc ? fromDocument(doc) : null;
}

/**
 * Find feedback by organization.
 */
export async function findByOrganization(organizationId: string): Promise<Feedback[]> {
	const coll = await getCollection();
	const docs = await coll
		.find({ organizationId: new ObjectId(organizationId) }, { sort: { _id: -1 } })
		.toArray();
	return docs.map(fromDocument);
}

/**
 * Find feedback by project.
 */
export async function findByProject(
	organizationId: string,
	projectId: string
): Promise<Feedback[]> {
	const coll = await getCollection();
	const docs = await coll
		.find(
			{ organizationId: new ObjectId(organizationId), projectId: new ObjectId(projectId) },
			{ sort: { _id: -1 } }
		)
		.toArray();
	return docs.map(fromDocument);
}

/**
 * Remove all feedback that belong to an organization.
 */
export async function removeAll(organizationId: string): Promise<void> {
	const coll = await getCollection();
	await coll.deleteMany({ organizationId: new ObjectId(organizationId) });
}
