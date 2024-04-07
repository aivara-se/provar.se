<script lang="ts">
	import { exploreWithValue } from '$lib/client/search';
	import type { Feedback, Organization } from '$lib/client/types';
	import { FilterIcon } from 'lucide-svelte';

	export let feedback: Feedback;
	export let organization: Organization;

	function filterByMetadata(key: string, val: string) {
		exploreWithValue(organization.slug, { meta: { [key]: val } });
	}
</script>

{#if Object.keys(feedback.feedbackMeta || {}).length}
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
				{#each Object.keys(feedback.feedbackMeta || {}).sort() as key}
					<tr>
						<th>{key}</th>
						<td>{feedback.feedbackMeta?.[key]}</td>
						<td class="text-right">
							<button
								class="btn btn-sm btn-ghost"
								on:click={() => filterByMetadata(key, feedback.feedbackMeta?.[key])}
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
