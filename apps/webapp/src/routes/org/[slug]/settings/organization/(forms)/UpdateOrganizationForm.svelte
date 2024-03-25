<script lang="ts">
	import { api } from '$lib/client/api';
	import { TextArea, TextInput } from '$lib/client/forms';
	import { toast } from '$lib/client/toast';
	import type { Organization } from '$lib/client/types';
	import { kebabCase } from 'lodash';
	import { superForm } from 'sveltekit-superforms';
	import { zod } from 'sveltekit-superforms/adapters';
	import { schema } from './UpdateOrganizationForm.schema';
	import { goto } from '$app/navigation';

	export let formData;
	export let organization: Organization;

	const { form, errors, constraints, enhance } = superForm(formData, {
		SPA: true,
		validators: zod(schema),
		onUpdate: async ({ form }) => {
			await api().Organization.updateDetails(organization.id, {
				name: form.data.name,
				slug: kebabCase(form.data.name),
				description: form.data.description
			});
			goto(`/org/${organization.slug}/settings/organization`);
		},
		onResult: ({ result }) => {
			if (result.type === 'failure') {
				toast('error', 'Failed to update organization details');
			} else if (result.type === 'success') {
				toast('success', 'Organization details updated successfully');
			}
		}
	});
</script>

<form method="POST" action="?/updateOrganization" class="flex flex-col gap-4" use:enhance>
	<TextInput
		name="name"
		label="Name"
		errors={$errors.name}
		constraints={$constraints.name}
		bind:value={$form.name}
	/>
	<TextArea
		name="description"
		label="Description"
		errors={$errors.description}
		constraints={$constraints.description}
		bind:value={$form.description}
	/>
	<div class="flex mt-4">
		<button class="btn">Update</button>
	</div>
</form>
