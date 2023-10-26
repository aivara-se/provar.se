import { ObjectId, type Collection, type MongoClient } from 'mongodb';
import { beforeAll, beforeEach, describe, expect, it } from 'vitest';
import { getMongoClient } from '../database';
import * as featureRepository from './feature.repository';

const unknownId = new ObjectId();
const featureId1 = new ObjectId();
const featureId2 = new ObjectId();
const featureId3 = new ObjectId();
const organizationId1 = new ObjectId();

describe('Feature Repository', () => {
	let mongoClient: MongoClient;
	let collection: Collection;

	beforeAll(async () => {
		mongoClient = await getMongoClient();
		collection = mongoClient.db().collection('features');
	});

	beforeEach(async () => {
		await collection.deleteMany({});
		await collection.insertMany([
			{ _id: featureId1, key: 'f1', organizationId: organizationId1 },
			{ _id: featureId2, key: 'f2', organizationId: organizationId1 },
			{ _id: featureId3, key: 'f3', organizationId: organizationId1 }
		]);
	});

	describe('findById', () => {
		it('should return null if there are no matching documents', async () => {
			const result = await featureRepository.findById(unknownId.toHexString());
			expect(result).toBeNull();
		});

		it('should return data if there is a matching document', async () => {
			const result = await featureRepository.findById(featureId1.toHexString());
			expect(result).toEqual(
				expect.objectContaining({
					id: featureId1.toHexString(),
					organizationId: organizationId1.toHexString()
				})
			);
		});
	});

	describe('findByKey', () => {
		it('should return null if there are no matching documents', async () => {
			const result = await featureRepository.findByKey('f0');
			expect(result).toBeNull();
		});

		it('should return data if there is a matching document', async () => {
			const result = await featureRepository.findByKey('f1');
			expect(result).toEqual(
				expect.objectContaining({
					id: featureId1.toHexString(),
					key: 'f1',
					organizationId: organizationId1.toHexString()
				})
			);
		});
	});

	describe('findByOrganization', () => {
		it('should return empty array if there are no matching documents', async () => {
			const result = await featureRepository.findByOrganization(unknownId.toHexString());
			expect(result).toEqual([]);
		});

		it('should return data if there are matching documents', async () => {
			const result = await featureRepository.findByOrganization(organizationId1.toHexString());
			expect(result).toEqual([
				expect.objectContaining({
					id: featureId1.toHexString(),
					organizationId: organizationId1.toHexString()
				}),
				expect.objectContaining({
					id: featureId2.toHexString(),
					organizationId: organizationId1.toHexString()
				}),
				expect.objectContaining({
					id: featureId3.toHexString(),
					organizationId: organizationId1.toHexString()
				})
			]);
		});
	});
});
