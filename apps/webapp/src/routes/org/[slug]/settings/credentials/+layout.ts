import { api } from '$lib/client/api';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async (event) => {
	const { organization } = await event.parent();
	const credentials = await api().Credential.list(organization.id);
	return { credentials };
};
