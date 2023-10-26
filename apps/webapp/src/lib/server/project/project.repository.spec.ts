import { ObjectId, type Collection, type MongoClient } from 'mongodb';
import { beforeAll, beforeEach, describe, expect, it } from 'vitest';
import { getMongoClient } from '../database';
import * as projectRepository from './project.repository';

const unknownId = new ObjectId();
const projectId1 = new ObjectId();
const projectId2 = new ObjectId();
const projectId3 = new ObjectId();
const organizationId1 = new ObjectId();

describe('Project Repository', () => {
	let mongoClient: MongoClient;
	let collection: Collection;

	beforeAll(async () => {
		mongoClient = await getMongoClient();
		collection = mongoClient.db().collection('projects');
	});

	beforeEach(async () => {
		await collection.deleteMany({});
		await collection.insertMany([
			{ _id: projectId1, organizationId: organizationId1 },
			{ _id: projectId2, organizationId: organizationId1 },
			{ _id: projectId3, organizationId: organizationId1 }
		]);
	});

	describe('findById', () => {
		it('should return null if there are no matching documents', async () => {
			const result = await projectRepository.findById(unknownId.toHexString());
			expect(result).toBeNull();
		});

		it('should return data if there is a matching document', async () => {
			const result = await projectRepository.findById(projectId1.toHexString());
			expect(result).toEqual(
				expect.objectContaining({
					id: projectId1.toHexString(),
					organizationId: organizationId1.toHexString()
				})
			);
		});
	});

	describe('findByOrganization', () => {
		it('should return empty array if there are no matching documents', async () => {
			const result = await projectRepository.findByOrganization(unknownId.toHexString());
			expect(result).toEqual([]);
		});

		it('should return data if there are matching documents', async () => {
			const result = await projectRepository.findByOrganization(organizationId1.toHexString());
			expect(result).toEqual([
				expect.objectContaining({
					id: projectId1.toHexString(),
					organizationId: organizationId1.toHexString()
				}),
				expect.objectContaining({
					id: projectId2.toHexString(),
					organizationId: organizationId1.toHexString()
				}),
				expect.objectContaining({
					id: projectId3.toHexString(),
					organizationId: organizationId1.toHexString()
				})
			]);
		});
	});
});
