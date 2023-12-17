import { getMongoClient } from '$lib/server/database';
import { ObjectId, type Collection, type MongoClient } from 'mongodb';
import { afterAll, beforeAll, beforeEach, describe, expect, it } from 'vitest';
import {
	addMember,
	create,
	findById,
	findByMember,
	remove,
	removeMember,
	update
} from './organization.repository';

const unknownId = new ObjectId();
const userId1 = new ObjectId();
const userId2 = new ObjectId();
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

	afterAll(async () => {
		// Clean up after all tests
		await mongoClient.close();
	});

	beforeEach(async () => {
		await collection.deleteMany({});
		await collection.insertMany([
			{ _id: organizationId1, members: [userId1] },
			{ _id: organizationId2, members: [userId1, userId2] },
			{ _id: organizationId3, members: [userId2] }
		]);
	});

	describe('create', () => {
		it('should create a new organization', async () => {
			const id = await create(userId1.toHexString(), {
				name: 'Test'
			});
			const result = await collection.findOne({ _id: new ObjectId(id) });
			expect(result).toEqual(
				expect.objectContaining({
					name: 'Test',
					members: [userId1]
				})
			);
		});
	});

	describe('update', () => {
		it('should update an existing organization', async () => {
			const id = await create(userId1.toHexString(), {
				name: 'Test'
			});
			const updateData = { name: 'Updated Test', description: 'Updated Desc' };
			await update(id, updateData);
			const result = await collection.findOne({ _id: new ObjectId(id) });
			expect(result).toEqual(expect.objectContaining(updateData));
		});
	});

	describe('remove', () => {
		it('should remove an existing organization', async () => {
			const id = await create(userId1.toHexString(), {
				name: 'Test'
			});
			await remove(id);
			const result = await collection.findOne({ _id: new ObjectId(id) });
			expect(result).toBeNull();
		});
	});

	describe('addMember', () => {
		it('should add a new member to an organization', async () => {
			const newUserId = new ObjectId().toHexString();
			await addMember(organizationId1.toHexString(), newUserId);
			const organization = await collection.findOne({ _id: organizationId1 });
			const members = organization?.members?.map((m: ObjectId) => m.toHexString());
			expect(members).toContain(newUserId);
			expect(members).toContain(userId1.toHexString());
		});

		it('should not throw if the user is already a member', async () => {
			await addMember(organizationId1.toHexString(), userId1.toHexString());
			const organization = await collection.findOne({ _id: organizationId1 });
			const members = organization?.members?.map((m: ObjectId) => m.toHexString());
			expect(members).toContain(userId1.toHexString());
		});
	});

	describe('removeMember', () => {
		it('should remove an existing member from an organization', async () => {
			await removeMember(organizationId2.toHexString(), userId2.toHexString());
			const organization = await collection.findOne({ _id: organizationId2 });
			const members = organization?.members?.map((m: ObjectId) => m.toHexString());
			expect(members).not.toContain(userId2.toHexString());
			expect(members).toContain(userId1.toHexString());
		});

		it('should not throw if the user is not a member (removed?)', async () => {
			const nonExistingUserId = new ObjectId().toHexString();
			await removeMember(organizationId1.toHexString(), nonExistingUserId);
			const organization = await collection.findOne({ _id: organizationId1 });
			const members = organization?.members?.map((m: ObjectId) => m.toHexString());
			expect(members).toContain(userId1.toHexString());
			expect(members).not.toContain(nonExistingUserId);
		});
	});

	describe('findById', () => {
		it('should return null if there are no matching documents', async () => {
			const result = await findById(unknownId.toHexString());
			expect(result).toBeNull();
		});

		it('should return data if there is a matching document', async () => {
			const result = await findById(organizationId1.toHexString());
			expect(result).toEqual(expect.objectContaining({ id: organizationId1.toHexString() }));
		});
	});

	describe('findByMember', () => {
		it('should return empty array if there are no matching documents', async () => {
			const result = await findByMember(unknownId.toHexString());
			expect(result).toEqual([]);
		});

		it('should return data if there are matching documents', async () => {
			const result = await findByMember(userId1.toHexString());
			expect(result).toEqual([
				expect.objectContaining({
					id: organizationId1.toHexString(),
					members: [userId1.toHexString()]
				}),
				expect.objectContaining({
					id: organizationId2.toHexString(),
					members: [userId1.toHexString(), userId2.toHexString()]
				})
			]);
		});
	});
});
