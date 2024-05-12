<script lang="ts">
	import { exploreWithValue } from '$lib/client/search';
	import type { Feedback, Organization } from '$lib/client/types';
	import { FilterIcon } from 'lucide-svelte';

	export let feedback: Feedback;
	export let organization: Organization;

	$: metadataKeys = Object.keys(feedback?.feedbackMeta || {}).sort();

	function filterByMetadata(key: string, val: string) {
		exploreWithValue(organization.slug, { meta: { [key]: val } });
	}
</script>

{#if Object.keys(feedback.feedbackMeta || {}).length}
	<section class="mt-8">
		<h3 class="text-lg font-semibold">Feedback Metadata</h3>

		<table class="hidden md:table mt-4">
			<thead>
				<tr>
					<th>Key</th>
					<th>Value</th>
					<th></th>
				</tr>
			</thead>
			<tbody>
				{#each metadataKeys as key}
					{#if feedback.feedbackMeta?.[key]}
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
					{/if}
				{/each}
			</tbody>
		</table>

		<div class="flex flex-col md:hidden mt-4">
			{#each metadataKeys as key}
				{#if feedback.feedbackMeta?.[key]}
					<div class="flex items-center justify-between border-b border-gray-800 py-4">
						<div class="flex-1">
							<p class="text-sm font-semibold">{key}</p>
							<p class="text-sm text-gray-600">{feedback.feedbackMeta?.[key]}</p>
						</div>
						<button
							class="btn btn-sm btn-ghost"
							on:click={() => filterByMetadata(key, feedback.feedbackMeta?.[key])}
						>
							<FilterIcon class="w-3 h-3" />
						</button>
					</div>
				{/if}
			{/each}
		</div>
	</section>
{/if}
