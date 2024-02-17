<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import RevokeMembershipBlock from './(components)/RevokeMembershipBlock.svelte';
	import route from './route.meta';

	export let data;

	$: member = data.member;
</script>

<DashLayout {route}>
	<div class="flex justify-center mt-8">
		<div class="avatar">
			<div class="w-16 rounded-lg">
				<img src={member.image} alt={member.name} />
			</div>
		</div>
	</div>

	<table class="table mt-8">
		<tbody>
			<tr>
				<th class="min-w-32">Name</th>
				<td>{member.name}</td>
			</tr>
			<tr>
				<th class="min-w-32">Email</th>
				<td>{member.email}</td>
			</tr>
		</tbody>
	</table>

	<RevokeMembershipBlock {member} />
</DashLayout>
