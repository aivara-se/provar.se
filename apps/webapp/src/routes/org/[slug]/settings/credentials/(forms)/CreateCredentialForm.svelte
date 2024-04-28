<script lang="ts">
	import { api } from '$lib/client/api';
	import { TextInput } from '$lib/client/forms';
	import { toast } from '$lib/client/toast';
	import type { Organization } from '$lib/client/types';
	import { superForm } from 'sveltekit-superforms';
	import { zod } from 'sveltekit-superforms/adapters';
	import { schema } from './CreateCredentialForm.schema';

	export let formData;
	export let organization: Organization;

	const { form, errors, constraints, enhance } = superForm(formData, {
		SPA: true,
		validators: zod(schema),
		onUpdate: async ({ form }) => {
			await api().Credential.create(organization.id, { name: form.data.name });
		},
		onResult: ({ result }) => {
			if (result.type === 'failure') {
				toast('error', 'Failed to create credential');
			} else if (result.type === 'success') {
				toast('success', 'Credential created successfully');
			}
		}
	});
</script>

<section class="mt-8 rounded-lg p-4 bg-gray-950 text-gray-200">
	<h3 class="flex items-center justify-between text-lg font-medium">Create Credential</h3>
	<p class="text-sm mt-2">Credentials are used to authenticate your application with the API.</p>
	<form method="POST" action="?/createCredential" class="flex flex-col gap-2 mt-2" use:enhance>
		<TextInput
			name="name"
			label="Name"
			className="border-none bg-white/10"
			errors={$errors.name}
			constraints={$constraints.name}
			bind:value={$form.name}
		/>
		<button class="btn btn-sm w-24 btn-neutral mt-4">Create</button>
	</form>
</section>
