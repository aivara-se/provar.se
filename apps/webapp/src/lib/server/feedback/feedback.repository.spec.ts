import { ObjectId, type Collection, type MongoClient } from 'mongodb';
import { beforeAll, describe, expect, it } from 'vitest';
import { getMongoClient } from '../database';
import { findById } from './feedback.repository';

describe('Feedback Repository', () => {
	let mongoClient: MongoClient;
	let collection: Collection;

	beforeAll(async () => {
		mongoClient = await getMongoClient();
		collection = mongoClient.db().collection('feedbacks');
	});

	describe('findById', () => {
		it('should return null if there are no matching documents', async () => {
			const testFeedbackId = new ObjectId();
			const result = await findById(testFeedbackId.toHexString());
			expect(result).toBeNull();
		});

		it('should return data if there is a matching document', async () => {
			// setup
			const testFeedbackId = new ObjectId();
			const testOrganizationId = new ObjectId();
			await collection.insertOne({
				_id: testFeedbackId,
				organizationId: testOrganizationId
			});
			// test
			const result = await findById(testFeedbackId.toHexString());
			expect(result).toEqual({
				id: testFeedbackId.toHexString(),
				organizationId: testOrganizationId.toHexString()
			});
		});
	});
});
