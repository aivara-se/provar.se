<script lang="ts">
	import { ChevronRightIcon } from 'lucide-svelte';
	import type { Credential } from '$lib/types';

	export let credentials: Credential[] = [];
</script>

{#if credentials.length}
	<table class="table">
		<thead>
			<tr>
				<th>Name</th>
				<th>Created at</th>
				<th>Last used</th>
				<th></th>
			</tr>
		</thead>
		<tbody>
			{#each credentials as credential}
				<tr>
					<td>{credential.name}</td>
					<td>{credential.createdAt.toLocaleDateString()}</td>
					<td>{credential.lastUsedAt?.toLocaleDateString()}</td>
					<td class="text-right">
						<a href="./credentials/{credential.id}" role="button" class="btn btn-sm btn-ghost">
							<ChevronRightIcon class="w-4 h-4" />
						</a>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
{/if}
