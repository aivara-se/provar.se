<script lang="ts">
	import { TextInput } from '$lib/client/forms';
	import { toast } from '$lib/client/toast';
	import { superForm } from 'sveltekit-superforms';

	export let formData;

	const { form, errors, constraints, enhance } = superForm(formData, {
		onResult: ({ result }) => {
			if (result.type === 'failure') {
				toast('error', 'Failed to invite member');
			} else if (result.type === 'success') {
				toast('success', 'User invited successfully');
			}
		}
	});
</script>

<section class="mt-8 rounded-lg p-4 bg-gray-100 text-gray-900 dark:bg-gray-950 dark:text-gray-200">
	<h3 class="flex items-center justify-between text-lg font-medium">Invite Member</h3>
	<p class="text-sm mt-2">
		Invite a user to join your organization. The user will receive an email from provar.se with a
		link to join your organization.
	</p>
	<form method="POST" action="?/inviteMember" class="flex flex-col gap-2 mt-2" use:enhance>
		<TextInput
			name="name"
			label="Name"
			className="border-none bg-black/10 dark:bg-white/10"
			errors={$errors.name}
			constraints={$constraints.name}
			bind:value={$form.name}
		/>
		<TextInput
			name="email"
			label="Email"
			className="border-none bg-black/10 dark:bg-white/10"
			errors={$errors.email}
			constraints={$constraints.email}
			bind:value={$form.email}
		/>
		<button class="btn btn-sm w-24 btn-neutral mt-4">Invite</button>
	</form>
</section>
