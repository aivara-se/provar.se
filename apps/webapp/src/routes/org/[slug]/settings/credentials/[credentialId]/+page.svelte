<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import { formatDateString } from '$lib/client/dates';
	import DeleteCredentialBlock from './(components)/DeleteCredentialBlock.svelte';
	import route from './route.meta';

	export let data;

	$: credential = data.credential;
</script>

<DashLayout {route} user={data.user} organizations={data.organizations}>
	<table class="table mt-4">
		<tbody>
			<tr>
				<th class="min-w-32">Name</th>
				<td>{credential.name}</td>
			</tr>
			<tr>
				<th class="min-w-32">Created at</th>
				<td>{formatDateString(credential.createdAt)}</td>
			</tr>
			<tr>
				<th class="min-w-32">Last used at</th>
				<td>{formatDateString(credential.lastUsedAt) || '-'}</td>
			</tr>
			<tr>
				<th class="min-w-32">Secret key</th>
				<td><code class="break-all">{credential.secret}</code></td>
			</tr>
		</tbody>
	</table>

	<div>
		<DeleteCredentialBlock organization={data.organization} {credential} />
	</div>
</DashLayout>
