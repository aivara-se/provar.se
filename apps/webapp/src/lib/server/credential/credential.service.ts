import { secureRandom } from '$lib/server/random';
import type { Credential } from '$lib/types';
import * as CredentialRepository from './credential.repository';

/**
 * The length of the credential key.
 */
const CREDENTIAL_KEY_LENGTH = 64;

/**
 * Returns the credential used for importing. Create if necessary.
 */
export async function getImportCredential(organizationId: string): Promise<Credential> {
	const credentials = await CredentialRepository.findByOrganization(organizationId);
	const existingCredential = credentials.find((cred) => cred.name === 'import');
	if (existingCredential) {
		return existingCredential;
	}
	const credentialId = await CredentialRepository.create({
		organizationId,
		name: 'import',
		key: createCredentialKey()
	});
	const createdCredential = await CredentialRepository.findById(organizationId, credentialId);
	if (!createdCredential) {
		throw new Error('Unexpected Error: failed to create credential');
	}
	return createdCredential;
}

/**
 * Generates a secure credential key.
 */
export function createCredentialKey() {
	return secureRandom(CREDENTIAL_KEY_LENGTH);
}
