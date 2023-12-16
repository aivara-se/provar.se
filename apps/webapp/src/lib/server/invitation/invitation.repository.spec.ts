import { ObjectId, type Collection, type MongoClient } from 'mongodb';
import { beforeAll, beforeEach, describe, expect, it } from 'vitest';
import { getMongoClient } from '../database';
import * as invitationRepository from './invitation.repository';

const unknownId = new ObjectId();
const invitationId1 = new ObjectId();
const invitationId2 = new ObjectId();
const invitationId3 = new ObjectId();
const organizationId1 = new ObjectId();

describe('Invitation Repository', () => {
	let mongoClient: MongoClient;
	let collection: Collection;

	beforeAll(async () => {
		mongoClient = await getMongoClient();
		collection = mongoClient.db().collection('invitations');
	});

	beforeEach(async () => {
		await collection.deleteMany({});
		await collection.insertMany([
			{
				_id: invitationId1,
				key: 'key-1',
				email: 'user1@gmail.com',
				organizationId: organizationId1
			},
			{
				_id: invitationId2,
				key: 'key-2',
				email: 'user2@gmail.com',
				organizationId: organizationId1
			},
			{
				_id: invitationId3,
				key: 'key-3',
				email: 'user3@gmail.com',
				organizationId: organizationId1
			}
		]);
	});

	describe('create', () => {
		it('should create a new invitation', async () => {
			await invitationRepository.create({
				key: 'key-0',
				email: 'user0@gmail.com',
				organizationId: organizationId1.toHexString()
			});
			const result = await collection.findOne({ key: 'key-0' });
			expect(result).toEqual(
				expect.objectContaining({
					key: 'key-0',
					email: 'user0@gmail.com',
					organizationId: organizationId1
				})
			);
		});
	});

	describe('remove', () => {
		it('should remove a invitation', async () => {
			await invitationRepository.remove(organizationId1.toHexString(), invitationId1.toHexString());
			const result = await collection.findOne({ _id: invitationId1 });
			expect(result).toBeNull();
		});
	});

	describe('findByKey', () => {
		it('should return null if there are no matching documents', async () => {
			const result = await invitationRepository.findByKey('key-0');
			expect(result).toBeNull();
		});

		it('should return data if there is a matching document', async () => {
			const result = await invitationRepository.findByKey('key-1');
			expect(result).toEqual(
				expect.objectContaining({
					key: 'key-1',
					organizationId: organizationId1.toHexString()
				})
			);
		});
	});

	describe('findByOrganization', () => {
		it('should return empty array if there are no matching documents', async () => {
			const result = await invitationRepository.findByOrganization(unknownId.toHexString());
			expect(result).toEqual([]);
		});

		it('should return data if there are matching documents', async () => {
			const result = await invitationRepository.findByOrganization(organizationId1.toHexString());
			expect(result).toEqual([
				expect.objectContaining({
					key: 'key-1',
					organizationId: organizationId1.toHexString()
				}),
				expect.objectContaining({
					key: 'key-2',
					organizationId: organizationId1.toHexString()
				}),
				expect.objectContaining({
					key: 'key-3',
					organizationId: organizationId1.toHexString()
				})
			]);
		});
	});
});
