import { getMongoClient } from '$lib/server/database';
import { ObjectId, type Collection, type MongoClient } from 'mongodb';
import { beforeAll, beforeEach, describe, expect, it } from 'vitest';
import { findById, findByOrganization, remove, removeAll } from './credential.repository';

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

	describe('remove', () => {
		it('should return success if attmpted to delete an unknown document', async () => {
			const count0 = (await collection.find({}).toArray()).length;
			await remove(unknownId.toHexString(), unknownId.toHexString());
			const count1 = (await collection.find({}).toArray()).length;
			expect(count1).toEqual(count0);
		});

		it('should return success after deleting a document', async () => {
			const count0 = (await collection.find({}).toArray()).length;
			await remove(organizationId1.toHexString(), credentialId1.toHexString());
			const count1 = (await collection.find({}).toArray()).length;
			expect(count1).toEqual(count0 - 1);
		});
	});

	describe('removeAll', () => {
		it('should return success if attempted to delete credentials of an unknown organization', async () => {
			const count0 = (await collection.find({}).toArray()).length;
			await removeAll(unknownId.toHexString());
			const count1 = (await collection.find({}).toArray()).length;
			expect(count1).toEqual(count0);
		});

		it('should return success after deleting all credentials of an organization', async () => {
			const count0 = (await collection.find({}).toArray()).length;
			await removeAll(organizationId1.toHexString());
			const count1 = (await collection.find({}).toArray()).length;
			expect(count1).toBeLessThan(count0);
		});
	});

	describe('findById', () => {
		it('should return null if there is no matching document', async () => {
			const result = await findById(organizationId1.toHexString(), unknownId.toHexString());
			expect(result).toBeNull();
		});

		it('should return the credential if there is a matching document', async () => {
			const result = await findById(organizationId1.toHexString(), credentialId1.toHexString());
			expect(result).toEqual(
				expect.objectContaining({
					id: credentialId1.toHexString()
				})
			);
		});
	});

	describe('findByOrganization', () => {
		it('should return empty array if there are no matching documents', async () => {
			const result = await findByOrganization(unknownId.toHexString());
			expect(result).toEqual([]);
		});

		it('should return data if there are matching documents', async () => {
			const result = await findByOrganization(organizationId1.toHexString());
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
