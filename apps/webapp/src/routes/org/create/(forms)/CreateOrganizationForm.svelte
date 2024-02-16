<script lang="ts">
	import { TextInput } from '$lib/client/forms';
	import { toast } from '$lib/client/toast';
	import { superForm } from 'sveltekit-superforms';

	export let formData;

	const { form, errors, constraints, enhance } = superForm(formData, {
		onResult: ({ result }) => {
			if (result.type === 'failure') {
				toast('error', 'Failed to create an organization');
			} else if (result.type === 'success') {
				toast('success', 'Organization created successfully');
			}
		}
	});
</script>

<form method="POST" action="?/createOrganization" class="flex flex-col gap-4" use:enhance>
	<TextInput
		name="name"
		label="Name"
		value={$form.name}
		errors={$errors.name}
		constraints={$constraints.name}
		className="input-bordered"
	/>
	<div class="flex mt-4">
		<button class="btn btn-block">Create organization</button>
	</div>
</form>
