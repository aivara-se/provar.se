import { env } from '$env/dynamic/private';
import { MongoClient } from 'mongodb';

let client: MongoClient;

export const getMongoConnection = async () => {
	if (!client) {
		client = new MongoClient(env.MONGODB_URL);
		await client.connect();
	}
	return client;
};
