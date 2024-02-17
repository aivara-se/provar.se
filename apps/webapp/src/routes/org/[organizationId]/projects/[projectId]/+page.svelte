<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import DeleteProjectBlock from './(components)/DeleteProjectBlock.svelte';
	import route from './route.meta';

	export let data;

	$: project = data.project;
</script>

<DashLayout {route}>
	<table class="table mt-4">
		<tbody>
			<tr>
				<th class="min-w-32">Name</th>
				<td>{project.name}</td>
			</tr>
			<tr>
				<th class="min-w-32">Created at</th>
				<td>{project.createdAt?.toLocaleDateString()}</td>
			</tr>
		</tbody>
	</table>

	<div>
		<DeleteProjectBlock {project} />
	</div>
</DashLayout>
