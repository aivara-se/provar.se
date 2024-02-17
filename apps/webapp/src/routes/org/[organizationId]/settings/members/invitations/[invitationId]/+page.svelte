<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import ResendInvitationBlock from './(components)/ResendInvitationBlock.svelte';
	import RevokeInvitationBlock from './(components)/RevokeInvitationBlock.svelte';
	import route from './route.meta';

	export let data;

	$: invitation = data.invitation;
</script>

<DashLayout {route}>
	<div class="flex justify-center mt-8">
		<div class="avatar">
			<div class="w-16 rounded-lg">
				<img src={invitation.image} alt={invitation.name} />
			</div>
		</div>
	</div>

	<table class="table mt-8">
		<tbody>
			<tr>
				<th class="min-w-32">Name</th>
				<td>{invitation.name}</td>
			</tr>
			<tr>
				<th class="min-w-32">Email</th>
				<td>{invitation.email}</td>
			</tr>
			<tr>
				<th class="min-w-32">Created at</th>
				<td>{invitation.createdAt.toLocaleDateString()}</td>
			</tr>
		</tbody>
	</table>

	<ResendInvitationBlock {invitation} />
	<RevokeInvitationBlock {invitation} />
</DashLayout>
