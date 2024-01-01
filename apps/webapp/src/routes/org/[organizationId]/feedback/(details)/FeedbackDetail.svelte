<script lang="ts">
	import type { Feedback } from '$lib/types';
	import { P } from 'flowbite-svelte';
	import FeedbackCNPS from './_components/FeedbackCNPS.svelte';
	import FeedbackCSAT from './_components/FeedbackCSAT.svelte';
	import FeedbackHead from './_components/FeedbackHead.svelte';
	import FeedbackMeta from './_components/FeedbackMeta.svelte';
	import FeedbackTags from './_components/FeedbackTags.svelte';
	import FeedbackText from './_components/FeedbackText.svelte';

	export let feedback: Feedback | null = null;
</script>

<section class="content">
	<div class="content-scroll">
		<div class="h-0">
			{#if !feedback}
				<section>
					<P class="text-gray-500 text-sm">Select a feedback from the list to see details</P>
				</section>
			{/if}

			{#if feedback}
				{#key feedback.id}
					<section>
						<FeedbackHead {feedback} />
						<FeedbackText {feedback} />
						<FeedbackTags {feedback} />

						{#if feedback.type === 'cnps'}
							<FeedbackCNPS {feedback} />
						{/if}

						{#if feedback.type === 'csat'}
							<FeedbackCSAT {feedback} />
						{/if}

						<FeedbackMeta {feedback} />
					</section>
				{/key}
			{/if}
		</div>
	</div>
</section>

<style>
	.content {
		display: flex;
		flex: 1;
		min-height: 0px;
		padding: 20px 0 0 20px;
		box-sizing: border-box;
	}

	.content-scroll {
		display: flex;
		flex: 1;
		flex-direction: column;
		overflow-x: hidden;
		overflow-y: auto;
	}

	.content-scroll::-webkit-scrollbar {
		display: none;
	}
</style>
