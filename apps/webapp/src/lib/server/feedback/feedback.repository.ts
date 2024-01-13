import { getMongoClient } from '$lib/server/database';
import type { SearchQuery } from '$lib/shared/search';
import {
	FeedbackType,
	type CNPSFeedbackData,
	type CSATFeedbackData,
	type Feedback,
	type FeedbackMeta,
	type FeedbackTags,
	type FeedbackUser,
	type TextFeedbackData,
	type TimeSeries,
	type DateRange
} from '$lib/types';
import { ObjectId, type Collection, type Filter } from 'mongodb';

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
	createdAt: Date;
	organizationId: ObjectId;
	projectId?: ObjectId;
	type: FeedbackType;
	meta?: FeedbackMeta;
	user?: FeedbackUser;
	tags?: FeedbackTags;
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
		createdAt: doc.createdAt,
		organizationId: doc.organizationId.toHexString(),
		projectId: doc.projectId?.toHexString(),
		type: doc.type,
		meta: doc.meta,
		user: doc.user,
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
 * Options for finding feedback.
 */
export interface FindOptions {
	page: number;
	limit: number;
	date?: DateRange;
	search?: SearchQuery;
}

/**
 * Find feedback by organization.
 */
export async function findByOrganization(
	organizationId: string,
	options: FindOptions
): Promise<Feedback[]> {
	const coll = await getCollection();
	const query: Filter<FeedbackDocument> = getFilterWithOptions(options, {
		organizationId: new ObjectId(organizationId)
	});
	const docs = await coll
		.find(query, {
			sort: { _id: -1 },
			limit: options.limit,
			skip: (options.page - 1) * options.limit
		})
		.toArray();
	return docs.map(fromDocument);
}

/**
 * Count feedback by organization.
 */
export async function countByOrganization(
	organizationId: string,
	options: FindOptions
): Promise<number> {
	const coll = await getCollection();
	const query: Filter<FeedbackDocument> = getFilterWithOptions(options, {
		organizationId: new ObjectId(organizationId)
	});
	return coll.countDocuments(query);
}

/**
 * Feedback count summary by date.
 */
export async function summarizeCountByOrganization(
	organizationId: string,
	options: FindOptions
): Promise<TimeSeries<number>> {
	const coll = await getCollection();
	const pipiline = [
		{
			$match: getFilterWithOptions(options, {
				organizationId: new ObjectId(organizationId)
			})
		},
		{
			$group: {
				_id: { $dateToString: { date: { $toDate: '$_id' }, format: '%Y-%m-%d' } },
				count: { $count: {} }
			}
		},
		{
			$sort: {
				_id: 1
			}
		}
	];
	const docs = await coll.aggregate(pipiline).toArray();
	return docs.map((doc) => ({ date: doc._id, value: doc.count }));
}

/**
 * Feedback CNPS summary by date.
 */
export async function summarizeCNPSByOrganization(
	organizationId: string,
	options: FindOptions
): Promise<TimeSeries<number>> {
	const coll = await getCollection();
	const pipiline = [
		{
			$match: getFilterWithOptions(options, {
				organizationId: new ObjectId(organizationId),
				type: FeedbackType.CNPS
			})
		},
		{
			$group: {
				_id: { $dateToString: { date: { $toDate: '$_id' }, format: '%Y-%m-%d' } },
				cnps: { $avg: '$data.cnps' }
			}
		},
		{
			$sort: {
				_id: 1
			}
		}
	];
	const docs = await coll.aggregate(pipiline).toArray();
	return docs.map((doc) => ({ date: doc._id, value: doc.cnps }));
}

/**
 * Feedback CSAT summary by date.
 */
export async function summarizeCSATByOrganization(
	organizationId: string,
	options: FindOptions
): Promise<TimeSeries<number>> {
	const coll = await getCollection();
	const pipiline = [
		{
			$match: getFilterWithOptions(options, {
				organizationId: new ObjectId(organizationId),
				type: FeedbackType.CSAT
			})
		},
		{
			$group: {
				_id: { $dateToString: { date: { $toDate: '$_id' }, format: '%Y-%m-%d' } },
				csat: { $avg: '$data.csat' }
			}
		},
		{
			$sort: {
				_id: 1
			}
		}
	];
	const docs = await coll.aggregate(pipiline).toArray();
	return docs.map((doc) => ({ date: doc._id, value: doc.csat }));
}

/**
 * Find feedback by project.
 */
export async function findByProject(
	organizationId: string,
	projectId: string,
	options: FindOptions
): Promise<Feedback[]> {
	const coll = await getCollection();
	const query: Filter<FeedbackDocument> = getFilterWithOptions(options, {
		organizationId: new ObjectId(organizationId),
		projectId: new ObjectId(projectId)
	});
	const docs = await coll
		.find(query, {
			sort: { _id: -1 },
			limit: options.limit,
			skip: (options.page - 1) * options.limit
		})
		.toArray();
	return docs.map(fromDocument);
}

/**
 * Count feedback by project.
 */
export async function countByProject(
	organizationId: string,
	projectId: string,
	options: FindOptions
): Promise<number> {
	const coll = await getCollection();
	const query: Filter<FeedbackDocument> = getFilterWithOptions(options, {
		organizationId: new ObjectId(organizationId),
		projectId: new ObjectId(projectId)
	});
	return coll.countDocuments(query);
}

/**
 * Remove all feedback that belong to an organization.
 */
export async function removeAll(organizationId: string): Promise<void> {
	const coll = await getCollection();
	await coll.deleteMany({ organizationId: new ObjectId(organizationId) });
}

/**
 * Get filter with additional options.
 */
function getFilterWithOptions(
	options: FindOptions,
	query: Filter<FeedbackDocument> = {}
): Filter<FeedbackDocument> {
	const updatedQuery: Filter<FeedbackDocument> = {};
	if (options.date) {
		updatedQuery._id = {
			$gte: ObjectId.createFromTime(Math.floor(options.date.from.getTime() / 1000)),
			$lte: ObjectId.createFromTime(Math.floor(options.date.to.getTime() / 1000))
		};
	}
	if (options.search?.text.length) {
		updatedQuery['data.text'] = { $regex: options.search.text.join('.*'), $options: 'si' };
	}
	if (options.search?.type?.length) {
		updatedQuery[`type`] = { $in: options.search.type };
	}
	if (options.search?.tags) {
		Object.entries(options.search.tags).forEach(([key, value]) => {
			updatedQuery[`tags.${key}`] = value;
		});
	}
	if (options.search?.meta) {
		Object.entries(options.search.meta).forEach(([key, value]) => {
			updatedQuery[`meta.${key}`] = value;
		});
	}
	return { ...updatedQuery, ...query };
}
