<script lang="ts">
	import { exploreWithValue } from '$lib/client/search';
	import type { Feedback } from '$lib/client/types';

	export let feedback: Feedback;

	function filterByType() {
		exploreWithValue(feedback.organizationId, { type: [feedback.feedbackType] });
	}

	function filterByTags(key: string, val: string) {
		exploreWithValue(feedback.organizationId, { tags: { [key]: val } });
	}
</script>

<section class="mt-3">
	<button class="badge text-xs mr-2 mb-2" on:click={filterByType}>
		type: {feedback.feedbackType}
	</button>

	{#each Object.keys(feedback.feedbackTags || {}).sort() as key}
		<button
			class="badge text-xs mr-2 mb-2"
			on:click={() => filterByTags(key, feedback.feedbackTags?.[key] || '')}
		>
			#{key}: {feedback.feedbackTags?.[key]}
		</button>
	{/each}
</section>
