import { ObjectId, type Collection, type MongoClient } from 'mongodb';
import { beforeAll, beforeEach, describe, expect, it } from 'vitest';
import { getMongoClient } from '../database';
import * as feedbackRepository from './feedback.repository';

const unknownId = new ObjectId();
const feedbackId1 = new ObjectId();
const feedbackId2 = new ObjectId();
const feedbackId3 = new ObjectId();
const projectId1 = new ObjectId();
const organizationId1 = new ObjectId();

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
			createDocument({ _id: feedbackId3, projectId: projectId1 })
		]);
	});

	describe('findById', () => {
		it('should return null if there are no matching documents', async () => {
			const result = await feedbackRepository.findById(unknownId.toHexString());
			expect(result).toBeNull();
		});

		it('should return data if there is a matching document', async () => {
			const result = await feedbackRepository.findById(feedbackId1.toHexString());
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
			const result = await feedbackRepository.findByOrganization(unknownId.toHexString());
			expect(result).toEqual([]);
		});

		it('should return data if there are matching documents', async () => {
			const result = await feedbackRepository.findByOrganization(organizationId1.toHexString());
			expect(result).toEqual([
				expect.objectContaining({
					id: feedbackId1.toHexString(),
					organizationId: organizationId1.toHexString()
				}),
				expect.objectContaining({
					id: feedbackId2.toHexString(),
					organizationId: organizationId1.toHexString()
				}),
				expect.objectContaining({
					id: feedbackId3.toHexString(),
					organizationId: organizationId1.toHexString()
				})
			]);
		});
	});

	describe('findByProject', () => {
		it('should return empty array if there are no matching documents', async () => {
			const result = await feedbackRepository.findByProject(
				unknownId.toHexString(),
				unknownId.toHexString()
			);
			expect(result).toEqual([]);
		});

		it('should return data if there are matching documents', async () => {
			const result = await feedbackRepository.findByProject(
				organizationId1.toHexString(),
				projectId1.toHexString()
			);
			expect(result).toEqual([
				expect.objectContaining({
					id: feedbackId2.toHexString(),
					organizationId: organizationId1.toHexString(),
					projectId: projectId1.toHexString()
				}),
				expect.objectContaining({
					id: feedbackId3.toHexString(),
					organizationId: organizationId1.toHexString(),
					projectId: projectId1.toHexString()
				})
			]);
		});
	});
});
