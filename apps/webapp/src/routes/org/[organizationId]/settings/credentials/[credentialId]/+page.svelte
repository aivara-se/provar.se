<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import DeleteCredentialBlock from './(components)/DeleteCredentialBlock.svelte';
	import route from './route.meta';

	export let data;

	$: credential = data.credential;
</script>

<DashLayout {route}>
	<table class="table mt-4">
		<tbody>
			<tr>
				<th class="min-w-32">Name</th>
				<td>{credential.name}</td>
			</tr>
			<tr>
				<th class="min-w-32">Created at</th>
				<td>{credential.createdAt?.toLocaleDateString()}</td>
			</tr>
			<tr>
				<th class="min-w-32">Last used at</th>
				<td>{credential.lastUsedAt?.toLocaleDateString()}</td>
			</tr>
			<tr>
				<th class="min-w-32">Secret key</th>
				<td><code class="break-all">{credential.key}</code></td>
			</tr>
		</tbody>
	</table>

	<div>
		<DeleteCredentialBlock {credential} />
	</div>
</DashLayout>
