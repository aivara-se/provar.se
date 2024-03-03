<script lang="ts">
	import { exploreWithValue } from '$lib/client/search';
	import type { Feedback } from '$lib/types';
	import { FilterIcon } from 'lucide-svelte';

	export let feedback: Feedback;

	function filterByMetadata(key: string, val: unknown) {
		exploreWithValue(feedback.organizationId, { meta: { [key]: val } });
	}
</script>

{#if Object.keys(feedback.meta || {}).length}
	<section class="mt-4">
		<h3 class="text-lg font-semibold">Feedback Metadata</h3>

		<table class="table mt-4">
			<thead>
				<tr>
					<th>Key</th>
					<th>Value</th>
					<th></th>
				</tr>
			</thead>

			<tbody>
				{#each Object.keys(feedback.meta || {}).sort() as key}
					<tr>
						<th>{key}</th>
						<td>{feedback.meta?.[key]}</td>
						<td class="text-right">
							<button
								class="btn btn-sm btn-ghost"
								on:click={() => filterByMetadata(key, feedback.meta?.[key])}
							>
								<FilterIcon class="w-3 h-3" />
							</button>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</section>
{/if}
