<script lang="ts">
	import { ArrowBigRightIcon } from 'lucide-svelte';
	import type { Credential } from '$lib/types';

	export let credentials: Credential[] = [];
</script>

<div>
	<table class="table">
		<thead>
			<tr>
				<th>Name</th>
				<th>Created At</th>
				<th>Last Used</th>
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
							<ArrowBigRightIcon class="w-4 h-4" />
						</a>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
