import { ObjectId, type Collection, type MongoClient } from 'mongodb';
import { beforeAll, beforeEach, describe, expect, it } from 'vitest';
import { getMongoClient } from '../database';
import * as credentialRepository from './credential.repository';

const unknownId = new ObjectId();
const credentialId1 = new ObjectId();
const credentialId2 = new ObjectId();
const credentialId3 = new ObjectId();
const organizationId1 = new ObjectId();

describe('Credential Repository', () => {
	let mongoClient: MongoClient;
	let collection: Collection;

	beforeAll(async () => {
		mongoClient = await getMongoClient();
		collection = mongoClient.db().collection('credentials');
	});

	beforeEach(async () => {
		await collection.deleteMany({});
		await collection.insertMany([
			{ _id: credentialId1, organizationId: organizationId1 },
			{ _id: credentialId2, organizationId: organizationId1 },
			{ _id: credentialId3, organizationId: organizationId1 }
		]);
	});

	describe('revoke', () => {
		it('should return success if attmpted to delete an unknown document', async () => {
			const count0 = (await collection.find({}).toArray()).length;
			await credentialRepository.revoke(unknownId.toHexString(), unknownId.toHexString());
			const count1 = (await collection.find({}).toArray()).length;
			expect(count1).toEqual(count0);
		});

		it('should return success after deleting a document', async () => {
			const count0 = (await collection.find({}).toArray()).length;
			await credentialRepository.revoke(organizationId1.toHexString(), credentialId1.toHexString());
			const count1 = (await collection.find({}).toArray()).length;
			expect(count1).toEqual(count0 - 1);
		});
	});

	describe('findByOrganization', () => {
		it('should return empty array if there are no matching documents', async () => {
			const result = await credentialRepository.findByOrganization(unknownId.toHexString());
			expect(result).toEqual([]);
		});

		it('should return data if there are matching documents', async () => {
			const result = await credentialRepository.findByOrganization(organizationId1.toHexString());
			expect(result).toEqual([
				expect.objectContaining({
					id: credentialId1.toHexString(),
					organizationId: organizationId1.toHexString()
				}),
				expect.objectContaining({
					id: credentialId2.toHexString(),
					organizationId: organizationId1.toHexString()
				}),
				expect.objectContaining({
					id: credentialId3.toHexString(),
					organizationId: organizationId1.toHexString()
				})
			]);
		});
	});
});
