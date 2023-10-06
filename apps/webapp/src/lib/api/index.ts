import { MOCK_FEEDBACKS_MAP, MOCK_ORGANIZATIONS } from './__mocks__';
import type { Feedback, Organization } from './types';

export async function getOrganizations(): Promise<Organization[]> {
	return MOCK_ORGANIZATIONS;
}

export async function getFeedbackByOrganizationId(organizationId: string): Promise<Feedback[]> {
	return MOCK_FEEDBACKS_MAP[organizationId];
}
