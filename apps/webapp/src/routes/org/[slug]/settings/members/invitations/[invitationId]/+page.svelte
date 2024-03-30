<script lang="ts">
	import { formatDateString } from '$lib/client/dates';
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
				<img src={invitation.avatar} alt={invitation.name} />
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
				<td>{formatDateString(invitation.createdAt)}</td>
			</tr>
		</tbody>
	</table>

	<ResendInvitationBlock invitation={data.invitation} />
	<RevokeInvitationBlock invitation={data.invitation} />
</DashLayout>
