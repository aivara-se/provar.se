<script lang="ts">
	import { type Feedback } from '$lib/types';
	import FeedbackItemCNPS from './FedbackItemCNPS.svelte';
	import FeedbackItemCSAT from './FedbackItemCSAT.svelte';
	import FeedbackItemText from './FedbackItemText.svelte';
	import FeedbackResultPagination from './FeedbackResultPagination.svelte';

	export let feedbacks = {
		count: 0,
		pages: 0,
		items: [] as Feedback[]
	};
</script>

{#if feedbacks.items.length}
	<section class="mt-4 flex flex-col gap-4">
		{#each feedbacks.items as feedback}
			<a
				href="./explore/{feedback.id}"
				class="bg-black/5 dark:bg-black/10 hover:bg-black/10 hover:dark:bg-black/20 transition-colors rounded-lg p-4 pt-6 md:pt-4 relative flex flex-col"
			>
				{#if feedback.type === 'csat'}
					<FeedbackItemCSAT {feedback} />
				{:else if feedback.type === 'cnps'}
					<FeedbackItemCNPS {feedback} />
				{:else}
					<FeedbackItemText {feedback} />
				{/if}
			</a>
		{/each}
	</section>
{/if}

{#if feedbacks.pages > 1}
	<FeedbackResultPagination pages={feedbacks.pages} />
{/if}
