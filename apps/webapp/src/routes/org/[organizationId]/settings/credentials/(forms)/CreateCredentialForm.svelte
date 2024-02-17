<script lang="ts">
	import { TextInput } from '$lib/client/forms';
	import { toast } from '$lib/client/toast';
	import { superForm } from 'sveltekit-superforms';

	export let formData;

	const { form, errors, constraints, enhance } = superForm(formData, {
		onResult: ({ result }) => {
			if (result.type === 'failure') {
				toast('error', 'Failed to create credential');
			} else if (result.type === 'success') {
				toast('success', 'Credential created successfully');
			}
		}
	});
</script>

<section class="mt-8 rounded-lg p-4 bg-gray-100 text-gray-900 dark:bg-gray-950 dark:text-gray-200">
	<h3 class="flex items-center justify-between text-lg font-medium">Create Credential</h3>
	<p class="text-sm mt-2">Credentials are used to authenticate your application with the API.</p>
	<form method="POST" action="?/createCredential" class="flex flex-col gap-2 mt-2" use:enhance>
		<TextInput
			name="name"
			label="Name"
			className="border-none bg-black/10 dark:bg-white/10"
			errors={$errors.name}
			constraints={$constraints.name}
			bind:value={$form.name}
		/>
		<button class="btn btn-sm w-24 btn-neutral mt-4">Create</button>
	</form>
</section>
