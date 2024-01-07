import { getMongoClient } from '$lib/server/database';
import { ObjectId, type Collection, type MongoClient } from 'mongodb';
import { beforeAll, beforeEach, describe, expect, it } from 'vitest';
import { findById, findByOrganization, findByProject, removeAll } from './feedback.repository';

const unknownId = new ObjectId();
const feedbackId1 = new ObjectId();
const feedbackId2 = new ObjectId();
const feedbackId3 = new ObjectId();
const feedbackId4 = new ObjectId();
const projectId1 = new ObjectId();
const organizationId1 = new ObjectId();
const organizationId2 = new ObjectId();

function createDocument(changes: object = {}): object {
	return {
		_id: new ObjectId(),
		createdAt: new Date(),
		organizationId: organizationId1,
		...changes
	};
}

describe('Feedback Repository', () => {
	let mongoClient: MongoClient;
	let collection: Collection;

	beforeAll(async () => {
		mongoClient = await getMongoClient();
		collection = mongoClient.db().collection('feedbacks');
	});

	beforeEach(async () => {
		await collection.deleteMany({});
		await collection.insertMany([
			createDocument({ _id: feedbackId1 }),
			createDocument({ _id: feedbackId2, projectId: projectId1 }),
			createDocument({ _id: feedbackId3, projectId: projectId1 }),
			createDocument({ _id: feedbackId4, organizationId: organizationId2 })
		]);
	});

	describe('findById', () => {
		it('should return null if there are no matching documents', async () => {
			const result = await findById(organizationId1.toHexString(), unknownId.toHexString());
			expect(result).toBeNull();
		});

		it('should return data if there is a matching document', async () => {
			const result = await findById(organizationId1.toHexString(), feedbackId1.toHexString());
			expect(result).toEqual(
				expect.objectContaining({
					id: feedbackId1.toHexString(),
					organizationId: organizationId1.toHexString()
				})
			);
		});
	});

	describe('findByOrganization', () => {
		it('should return empty array if there are no matching documents', async () => {
			const options = { page: 1, limit: 10 };
			const result = await findByOrganization(unknownId.toHexString(), options);
			expect(result).toEqual([]);
		});

		it('should return data if there are matching documents', async () => {
			const options = { page: 1, limit: 10 };
			const result = await findByOrganization(organizationId1.toHexString(), options);
			expect(result).toEqual([
				expect.objectContaining({
					id: feedbackId3.toHexString(),
					organizationId: organizationId1.toHexString()
				}),
				expect.objectContaining({
					id: feedbackId2.toHexString(),
					organizationId: organizationId1.toHexString()
				}),
				expect.objectContaining({
					id: feedbackId1.toHexString(),
					organizationId: organizationId1.toHexString()
				})
			]);
		});
	});

	describe('findByProject', () => {
		it('should return empty array if there are no matching documents', async () => {
			const options = { page: 1, limit: 10 };
			const result = await findByProject(unknownId.toHexString(), unknownId.toHexString(), options);
			expect(result).toEqual([]);
		});

		it('should return data if there are matching documents', async () => {
			const options = { page: 1, limit: 10 };
			const result = await findByProject(
				organizationId1.toHexString(),
				projectId1.toHexString(),
				options
			);
			expect(result).toEqual([
				expect.objectContaining({
					id: feedbackId3.toHexString(),
					organizationId: organizationId1.toHexString(),
					projectId: projectId1.toHexString()
				}),
				expect.objectContaining({
					id: feedbackId2.toHexString(),
					organizationId: organizationId1.toHexString(),
					projectId: projectId1.toHexString()
				})
			]);
		});
	});

	describe('removeAll', () => {
		it('should remove all feedback for the specified organization', async () => {
			await removeAll(organizationId1.toHexString());
			const org1Feedback = await collection.find({ organizationId: organizationId1 }).toArray();
			expect(org1Feedback).toEqual([]);
			const org2Feedback = await collection.find({ organizationId: organizationId2 }).toArray();
			expect(org2Feedback).toEqual([expect.objectContaining({ _id: feedbackId4 })]);
		});
	});
});
