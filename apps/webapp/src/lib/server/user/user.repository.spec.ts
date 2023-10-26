import { ObjectId, type Collection, type MongoClient } from 'mongodb';
import { beforeAll, beforeEach, describe, expect, it } from 'vitest';
import { getMongoClient } from '../database';
import * as userRepository from './user.repository';

const unknownId = new ObjectId();
const userId1 = new ObjectId();
const userId2 = new ObjectId();
const userId3 = new ObjectId();

describe('User Repository', () => {
	let mongoClient: MongoClient;
	let collection: Collection;

	beforeAll(async () => {
		mongoClient = await getMongoClient();
		collection = mongoClient.db().collection('users');
	});

	beforeEach(async () => {
		await collection.deleteMany({});
		await collection.insertMany([
			{ _id: userId1, email: 'f1@gmail.com' },
			{ _id: userId2, email: 'f2@gmail.com' },
			{ _id: userId3, email: 'f3@gmail.com' }
		]);
	});

	describe('findById', () => {
		it('should return null if there are no matching documents', async () => {
			const result = await userRepository.findById(unknownId.toHexString());
			expect(result).toBeNull();
		});

		it('should return data if there is a matching document', async () => {
			const result = await userRepository.findById(userId1.toHexString());
			expect(result).toEqual(
				expect.objectContaining({
					id: userId1.toHexString(),
					email: 'f1@gmail.com'
				})
			);
		});
	});

	describe('findByEmail', () => {
		it('should return null if there are no matching documents', async () => {
			const result = await userRepository.findByEmail('f0@gmail.com');
			expect(result).toBeNull();
		});

		it('should return data if there is a matching document', async () => {
			const result = await userRepository.findByEmail('f1@gmail.com');
			expect(result).toEqual(
				expect.objectContaining({
					id: userId1.toHexString(),
					email: 'f1@gmail.com'
				})
			);
		});
	});
});
