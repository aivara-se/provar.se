import { env } from '$env/dynamic/private';
import { MongoClient } from 'mongodb';

let client: MongoClient;

export const getMongoClient = async () => {
	if (!client && env.MONGODB_URL) {
		client = new MongoClient(env.MONGODB_URL);
		await client.connect();
	}
	return client;
};
