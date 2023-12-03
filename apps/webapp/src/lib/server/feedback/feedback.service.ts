import { env } from '$env/dynamic/private';
import { Storage } from '@google-cloud/storage';

/**
 * The expiration time for a signed URL in ms.
 * Right now it is set to 15 minutes.
 */
const IMPORT_URL_EXPIRATION = 15 * 60 * 1000;

/**
 * Returns the Google Cloud Storage instance.
 */
function getCloudStorage() {
	if (!env.GOOGLE_APPLICATION_CREDENTIALS) {
		return new Storage();
	}
	return new Storage({
		keyFilename: env.GOOGLE_APPLICATION_CREDENTIALS
	});
}

/**
 * The bucket where the feedback files to import are stored.
 */
function getImportBucket() {
	if (!env.GCS_IMPORT_BUCKET) {
		throw new Error('GCS_IMPORT_BUCKET is not defined');
	}
	return getCloudStorage().bucket(env.GCS_IMPORT_BUCKET);
}

/**
 * Returns the current date in UTC format.
 */
function getUTCDate(date = new Date()) {
	const year = date.getUTCFullYear();
	const month = String(date.getUTCMonth() + 1).padStart(2, '0');
	const day = String(date.getUTCDate()).padStart(2, '0');
	return `${year}-${month}-${day}`;
}

/**
 * Creates a signed URL for uploading a CSV file with feedbacks.
 */
export async function createUploadUrl(organizationId: string) {
	const fileName = `${getUTCDate()}-${secureRandom(8)}`;
	const filePath = `${organizationId}/${fileName}.csv`;
	const [signedUrl] = await getImportBucket()
		.file(filePath)
		.getSignedUrl({
			version: 'v4',
			action: 'write',
			expires: Date.now() + IMPORT_URL_EXPIRATION,
			contentType: 'text/csv'
		});
	return { signedUrl, fileName };
}

/**
 * Creates a signed URL for downloading a CSV file with feedbacks.
 */
export async function createReadableUrl(organizationId: string, fileName: string) {
	const filePath = `${organizationId}/${fileName}.csv`;
	const [signedUrl] = await getImportBucket()
		.file(filePath)
		.getSignedUrl({
			version: 'v4',
			action: 'read',
			expires: Date.now() + IMPORT_URL_EXPIRATION
		});
	return { signedUrl };
}

/**
 * Generates a secure random string.
 */
function secureRandom(count: number): string {
	const array = new Uint8Array(count);
	crypto.getRandomValues(array);
	return Buffer.from(array).toString('base64url');
}
