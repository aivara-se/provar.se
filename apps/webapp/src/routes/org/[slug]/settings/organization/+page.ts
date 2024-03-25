import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import type { PageLoad } from './$types';
import { schema } from './(forms)/UpdateOrganizationForm.schema';

export const load: PageLoad = async (event) => {
	const { organization } = await event.parent();
	const formData = { name: organization.name, description: organization.description };
	const form = await superValidate(formData, zod(schema));
	return { UpdateOrganizationForm: form };
};
