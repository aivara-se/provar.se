<script lang="ts">
	import { superForm } from 'sveltekit-superforms';
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

<form method="POST" action="?/updateOrganization" class="flex flex-col" use:enhance>
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Name</span>
		</div>
		<input
			type="text"
			name="name"
			class="input {$errors.name ? 'input-error' : ''}"
			aria-invalid={$errors.name ? 'true' : undefined}
			bind:value={$form.name}
			{...$constraints.name}
		/>
		{#if $errors.name}
			<div class="label">
				<span class="label-text text-xs text-red-700 dark:text-red-400">
					{$errors.name}
				</span>
			</div>
		{/if}
	</label>

	<label class="form-control mt-4">
		<div class="label"><span class="label-text">Description</span></div>
		<textarea
			name="description"
			class="textarea resize-none h-32 {$errors.description ? 'textarea-error' : ''}"
			aria-invalid={$errors.description ? 'true' : undefined}
			bind:value={$form.description}
			{...$constraints.description}
		/>
		{#if $errors.description}
			<div class="label">
				<span class="label-text text-xs text-red-700 dark:text-red-400">
					{$errors.description}
				</span>
			</div>
		{/if}
	</label>

	<div class="flex mt-8">
		<button class="btn">Update organization details</button>
	</div>
</form>
