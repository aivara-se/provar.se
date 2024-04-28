<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/client/api';
	import { ConfirmModal } from '$lib/client/forms';
	import type { Organization, User } from '$lib/client/types';
	import { AlertCircleIcon } from 'lucide-svelte';

	export let member: User;
	export let organization: Organization;

	let isOpen = false;

	async function revokeMembership() {
		await api().Organization.removeMember(organization.id, member.id);
		goto(`org/${organization.slug}/settings/members`);
	}
</script>

<section class="rounded-xl p-4 md:p-8 bg-black/20 text-gray-200 shadow mt-8">
	<h3 class="flex items-center justify-between text-lg font-medium">
		Remove membership!
		<AlertCircleIcon class="w-5 h-5" />
	</h3>
	<p class="text-sm mt-2">
		Removing membership will remove {member?.name} from the organization. The user will no longer be
		able to access feedback data.
	</p>
	<button class="btn btn-sm mt-4" on:click={() => (isOpen = true)}>
		Remove {member?.name}
	</button>
	<ConfirmModal bind:isOpen action={revokeMembership} submitText="Remove {member?.name}">
		<AlertCircleIcon class="w-8 h-8" />
		<p class="py-4 text-center">
			Are you sure you want to remove {member?.name}?
		</p>
	</ConfirmModal>
</section>
