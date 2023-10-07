import { getOrganizations, getFeedbackByOrganizationId } from '$lib/api';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
	const organizations = await getOrganizations();
	const feedback = await getFeedbackByOrganizationId(params.orgID);
	return { organizations, feedback };
};
