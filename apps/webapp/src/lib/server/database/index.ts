import { env } from '$env/dynamic/private';
import { MongoClient } from 'mongodb';

let client: MongoClient;

/**
 * Get the MongoDB client.
 */
export const getMongoClient = async () => {
	if (client) {
		return client;
	}
	if (!env.MONGODB_URL) {
		throw new Error('MONGODB_URL is not set');
	}
	client = new MongoClient(env.MONGODB_URL);
	await client.connect();
	return client;
};
