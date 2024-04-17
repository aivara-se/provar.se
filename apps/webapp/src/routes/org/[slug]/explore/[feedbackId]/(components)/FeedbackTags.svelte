<script lang="ts">
	import { exploreWithValue } from '$lib/client/search';
	import type { Feedback, Organization } from '$lib/client/types';

	export let feedback: Feedback;
	export let organization: Organization;

	function filterByType() {
		exploreWithValue(organization.slug, { type: [feedback.feedbackType] });
	}

	function filterByTags(key: string, val: string) {
		exploreWithValue(organization.slug, { tags: { [key]: val } });
	}
</script>

<section class="mt-4">
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
