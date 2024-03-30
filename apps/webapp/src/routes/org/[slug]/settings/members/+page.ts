import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import type { PageLoad } from './$types';
import { schema } from './(forms)/InviteMemberForm.schema';

export const load: PageLoad = async () => {
	const form = await superValidate(zod(schema));
	return { InviteMemberForm: form };
};
