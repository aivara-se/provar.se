<script lang="ts">
	import { ChevronRightIcon } from 'lucide-svelte';
	import type { Project } from '$lib/types';

	export let projects: Project[] = [];
</script>

{#if projects.length}
	<table class="table">
		<thead>
			<tr>
				<th>Name</th>
				<th>Created at</th>
				<th></th>
			</tr>
		</thead>
		<tbody>
			{#each projects as credential}
				<tr>
					<td>{credential.name}</td>
					<td>{credential.createdAt.toLocaleDateString()}</td>
					<td class="text-right">
						<a href="./projects/{credential.id}" role="button" class="btn btn-sm btn-ghost">
							<ChevronRightIcon class="w-4 h-4" />
						</a>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
{/if}
