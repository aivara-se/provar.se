<script lang="ts">
	import { superForm } from 'sveltekit-superforms';
	import { TextInput, TextArea } from '$lib/client/forms';
	import { toast } from '$lib/client/toast';

	export let formData;

	const { form, errors, constraints, enhance } = superForm(formData, {
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
