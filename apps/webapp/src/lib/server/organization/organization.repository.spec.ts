import { ObjectId, type Collection, type MongoClient } from 'mongodb';
import { beforeAll, beforeEach, describe, expect, it } from 'vitest';
import { getMongoClient } from '../database';
import * as organizationRepository from './organization.repository';

const unknownId = new ObjectId();
const userId1 = new ObjectId();
const organizationId1 = new ObjectId();
const organizationId2 = new ObjectId();
const organizationId3 = new ObjectId();

describe('Organization Repository', () => {
	let mongoClient: MongoClient;
	let collection: Collection;

	beforeAll(async () => {
		mongoClient = await getMongoClient();
		collection = mongoClient.db().collection('organizations');
	});

	beforeEach(async () => {
		await collection.deleteMany({});
		await collection.insertMany([
			{ _id: organizationId1, members: [userId1] },
			{ _id: organizationId2, members: [userId1] },
			{ _id: organizationId3, members: [] }
		]);
	});

	describe('create', () => {
		it('should create a new organization', async () => {
			await organizationRepository.create(userId1.toHexString(), {
				slug: 'test',
				name: 'Test',
				prod: false
			});
			const result = await collection.findOne({ slug: 'test' });
			expect(result).toEqual(
				expect.objectContaining({
					slug: 'test',
					name: 'Test',
					prod: false,
					members: [userId1]
				})
			);
		});
	});

	describe('findById', () => {
		it('should return null if there are no matching documents', async () => {
			const result = await organizationRepository.findById(unknownId.toHexString());
			expect(result).toBeNull();
		});

		it('should return data if there is a matching document', async () => {
			const result = await organizationRepository.findById(organizationId1.toHexString());
			expect(result).toEqual(expect.objectContaining({ id: organizationId1.toHexString() }));
		});
	});

	describe('findByMember', () => {
		it('should return empty array if there are no matching documents', async () => {
			const result = await organizationRepository.findByMember(unknownId.toHexString());
			expect(result).toEqual([]);
		});

		it('should return data if there are matching documents', async () => {
			const result = await organizationRepository.findByMember(userId1.toHexString());
			expect(result).toEqual([
				expect.objectContaining({
					id: organizationId1.toHexString(),
					members: [userId1.toHexString()]
				}),
				expect.objectContaining({
					id: organizationId2.toHexString(),
					members: [userId1.toHexString()]
				})
			]);
		});
	});
});
