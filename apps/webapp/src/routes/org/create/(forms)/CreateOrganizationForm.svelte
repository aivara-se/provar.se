<script lang="ts">
	import { getApi } from '$lib/client/api';
	import { TextInput } from '$lib/client/forms';
	import { kebabCase } from 'lodash';
	import { superForm } from 'sveltekit-superforms';
	import { zod } from 'sveltekit-superforms/adapters';
	import { schema } from './CreateOrganizationForm.schema';
	import { goto } from '$app/navigation';

	export let formData;

	const { form, errors, constraints, enhance } = superForm(formData, {
		SPA: true,
		validators: zod(schema),
		onUpdate: async ({ form }) => {
			const response = await getApi().Organization.create({
				name: form.data.name,
				slug: kebabCase(form.data.name)
			});
			goto(`/org/${response.slug}`);
		}
	});
</script>

<form method="POST" class="flex flex-col gap-4" use:enhance>
	<TextInput
		name="name"
		label="Name"
		bind:value={$form.name}
		errors={$errors.name}
		constraints={$constraints.name}
		className="input-bordered"
	/>
	<div class="flex mt-4">
		<button class="btn btn-block">Create organization</button>
	</div>
</form>
