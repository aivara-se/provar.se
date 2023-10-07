import type { Feedback, Organization } from '../types';

export const MOCK_ORGANIZATIONS: Organization[] = [
	{
		id: '2irh8ngfxxj',
		slug: 'wikimedia',
		name: 'Wikimedia',
		environment: 'production',
		members: []
	},
	{
		id: 'doh290140wd',
		slug: 'wikimedia-dev',
		name: 'Wikimedia - dev',
		environment: 'staging',
		members: []
	}
];

export const MOCK_FEEDBACKS_MAP: { [organizationId: string]: Feedback[] } = {
	'2irh8ngfxxj': [
		{
			id: '0ddv1yk92uy',
			organizationId: '2irh8ngfxxj',
			createdAt: 1696655680582,
			projectId: null,
			features: { 'feature-a': true },
			metadata: { userId: 'user-a' },
			formdata: {
				message: 'This is a test feedback message'
			}
		},
		{
			id: 'nmf4u4jsa64',
			organizationId: '2irh8ngfxxj',
			createdAt: 1696655673337,
			projectId: null,
			features: { 'feature-a': true, 'feature-b': true },
			metadata: { userId: 'user-b' },
			formdata: {
				message: 'This is another test feedback message'
			}
		}
	],
	doh290140wd: [
		{
			id: 'a3uaewtskhu',
			organizationId: 'doh290140wd',
			createdAt: 1696655665346,
			projectId: null,
			features: { 'feature-a': true, 'feature-b': true },
			metadata: { userId: 'user-c' },
			formdata: {
				message: 'One more test feedback message'
			}
		}
	]
};
