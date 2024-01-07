<script lang="ts">
	import type { FeedbackSummary } from '$lib/types';
	import CNPSSummary from './_components/CNPSSummary.svelte';
	import CSATSummary from './_components/CSATSummary.svelte';
	import CountSummary from './_components/CountSummary.svelte';

	export let summary: FeedbackSummary;
</script>

<section class="content">
	<div class="content-scroll">
		<main class="flex justify-between summary gap-4">
			{#key summary}
				<CNPSSummary values={summary.cnps} />
				<CSATSummary values={summary.csat} />
				<CountSummary values={summary.count} />
			{/key}
		</main>
	</div>
</section>

<style>
	.content {
		display: flex;
		flex: 1;
		min-height: 0px;
		padding: 10px 0 0 20px;
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
