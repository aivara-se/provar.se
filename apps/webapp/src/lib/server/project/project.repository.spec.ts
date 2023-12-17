import { getMongoClient } from '$lib/server/database';
import { ObjectId, type Collection, type MongoClient } from 'mongodb';
import { beforeAll, beforeEach, describe, expect, it } from 'vitest';
import {
	create,
	findById,
	findByOrganization,
	remove,
	removeAll,
	update
} from './project.repository';

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
			{ _id: projectId1, organizationId: new ObjectId(organizationId1) },
			{ _id: projectId2, organizationId: new ObjectId(organizationId1) },
			{ _id: projectId3, organizationId: new ObjectId(organizationId1) }
		]);
	});

	describe('create', () => {
		it('should create a new project', async () => {
			const projectId = await create({
				name: 'Test Project',
				organizationId: organizationId1.toHexString()
			});
			const result = await collection.findOne({ _id: new ObjectId(projectId) });
			expect(result).toEqual(
				expect.objectContaining({
					_id: new ObjectId(projectId),
					name: 'Test Project',
					organizationId: new ObjectId(organizationId1)
				})
			);
		});
	});

	describe('update', () => {
		it('should update an existing project', async () => {
			const projectId = await create({
				name: 'oldName',
				organizationId: organizationId1.toHexString()
			});
			const organizationId = organizationId1.toHexString();
			const newValues = { name: 'newName' };
			await update(organizationId, projectId, newValues);
			const result = await collection.findOne({ _id: new ObjectId(projectId) });
			expect(result).toEqual(
				expect.objectContaining({
					name: 'newName',
					organizationId: organizationId1
				})
			);
		});
	});

	describe('remove', () => {
		it('should remove a project', async () => {
			const projectId = await create({
				name: 'Test',
				organizationId: organizationId1.toHexString()
			});
			await remove(organizationId1.toHexString(), projectId);
			const result = await collection.findOne({ _id: new ObjectId(projectId) });
			expect(result).toBeNull();
		});
	});

	describe('removeAll', () => {
		it('should remove all projects of an organization', async () => {
			await create({
				name: 'Test1',
				organizationId: organizationId1.toHexString()
			});
			await create({
				name: 'Test2',
				organizationId: organizationId1.toHexString()
			});
			await removeAll(organizationId1.toHexString());
			const result = await collection.find({ organizationId: organizationId1 }).toArray();
			expect(result.length).toBe(0);
		});
	});

	describe('findById', () => {
		it('should return null if there are no matching documents', async () => {
			const result = await findById(unknownId.toHexString(), unknownId.toHexString());
			expect(result).toBeNull();
		});

		it('should return data if there is a matching document', async () => {
			const result = await findById(organizationId1.toHexString(), projectId1.toHexString());
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
			const result = await findByOrganization(unknownId.toHexString());
			expect(result).toEqual([]);
		});

		it('should return data if there are matching documents', async () => {
			const result = await findByOrganization(organizationId1.toHexString());
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
