import { defineRoute } from '$lib/client/routes';
import { LineChartIcon } from 'lucide-svelte';

export default defineRoute({
	id: 'org-feedback',
	sidebar: { weight: 10, icon: LineChartIcon, mobile: true },
	getName: () => 'Feedback',
	getPath: (params) => `/org/${params.organizationId}/feedback`,
	isActive: (path, params) => path.startsWith(`/org/${params.organizationId}/feedback`)
});
