import { MONGODB_URL } from '$env/static/private';
import { MongoClient } from 'mongodb';

const client = new MongoClient(MONGODB_URL);
const clientPromise = client.connect();

export const getMongoConnection = () => {
	return clientPromise;
};
