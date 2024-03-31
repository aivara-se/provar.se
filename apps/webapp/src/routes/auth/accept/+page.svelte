<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { api } from '$lib/client/api';
	import { AuthLayout } from '$lib/client/layout';

	export let data;

	$: organizationId = $page.url.searchParams.get('organizationId');
	$: invitationId = $page.url.searchParams.get('invitationId');
	$: secret = $page.url.searchParams.get('secret');
	$: hasRequiredParams = Boolean(organizationId && invitationId && secret);

	async function acceptInvitation() {
		await api().Invitation.accept(organizationId!, invitationId!, { secret: secret! });
		goto('/org');
	}
</script>

<AuthLayout>
	{#if hasRequiredParams}
		<div class="flex flex-col prose text-center">
			<h2>Join {data.invitedOrganization?.name}'s Workspace</h2>
			<p>
				You've received an invitation to join [Organization Name]'s workspace. Gain access to shared
				resources and collaborate seamlessly.
			</p>
			<p>
				<button class="btn btn-sm btn-outline" on:click={acceptInvitation}>Accept invitation</button
				>
			</p>
		</div>
	{:else}
		<div class="flex flex-col prose text-center">
			<h2>Unable to accept invitation</h2>
			<p>Please make sure the invitation link is valid.</p>
		</div>
	{/if}
</AuthLayout>
