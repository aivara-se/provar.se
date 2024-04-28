<script lang="ts">
	import { api } from '$lib/client/api';
	import { TextInput } from '$lib/client/forms';
	import { toast } from '$lib/client/toast';
	import { superForm } from 'sveltekit-superforms';
	import { zod } from 'sveltekit-superforms/adapters';
	import { schema } from './InviteMemberForm.schema';
	import type { Organization } from '$lib/client/types';

	export let formData;
	export let organization: Organization;

	const { form, errors, constraints, enhance } = superForm(formData, {
		SPA: true,
		validators: zod(schema),
		onUpdate: async ({ form }) => {
			await api().Invitation.create(organization.id, {
				name: form.data.name,
				email: form.data.email
			});
		},
		onResult: ({ result }) => {
			if (result.type === 'failure') {
				toast('error', 'Failed to invite member');
			} else if (result.type === 'success') {
				toast('success', 'User invited successfully');
			}
		}
	});
</script>

<section class="rounded-xl p-4 md:p-8 bg-black/20 text-gray-200 shadow mt-8">
	<h3 class="flex items-center justify-between text-lg font-medium">Invite Member</h3>
	<p class="text-sm mt-2">
		Invite a user to join your organization. The user will receive an email from provar.se with a
		link to join your organization.
	</p>
	<form method="POST" action="?/inviteMember" class="flex flex-col gap-2 mt-2" use:enhance>
		<TextInput
			name="name"
			label="Name"
			errors={$errors.name}
			constraints={$constraints.name}
			bind:value={$form.name}
		/>
		<TextInput
			name="email"
			label="Email"
			errors={$errors.email}
			constraints={$constraints.email}
			bind:value={$form.email}
		/>
		<button class="btn btn-sm w-24 mt-4">Invite</button>
	</form>
</section>
