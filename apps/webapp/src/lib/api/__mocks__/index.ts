import type { Feedback, Organization } from '../types';

export const MOCK_ORGANIZATIONS: Organization[] = [
	{ id: 'org-a', name: 'Wikipedia', environment: 'production' },
	{ id: 'org-b', name: 'Wikipedia - develop', environment: 'staging' }
];

export const MOCK_FEEDBACKS_MAP: { [organizationId: string]: Feedback[] } = {
	'org-a': [
		{
			id: 'feedback-a1',
			organizationId: 'org-a',
			createdAt: Date.now(),
			projectId: null,
			features: { 'feature-a': true },
			metadata: { userId: 'user-a' },
			formdata: {
				message: 'This is a test feedback message'
			}
		},
		{
			id: 'feedback-a2',
			organizationId: 'org-a',
			createdAt: Date.now(),
			projectId: null,
			features: { 'feature-a': true, 'feature-b': true },
			metadata: { userId: 'user-b' },
			formdata: {
				message: 'This is another test feedback message'
			}
		}
	],
	'org-b': [
		{
			id: 'feedback-b1',
			organizationId: 'org-b',
			createdAt: Date.now(),
			projectId: null,
			features: { 'feature-a': true, 'feature-b': true },
			metadata: { userId: 'user-c' },
			formdata: {
				message: 'One more test feedback message'
			}
		}
	]
};
