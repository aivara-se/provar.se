<script lang="ts">
	import type { Feedback } from '$lib/types';
	import { P } from 'flowbite-svelte';
	import FeedbackListItemCNPS from './_components/FeedbackListItemCNPS.svelte';
	import FeedbackListItemCSAT from './_components/FeedbackListItemCSAT.svelte';
	import FeedbackListItemText from './_components/FeedbackListItemText.svelte';
	import FeedbackListPagination from './_components/FeedbackListPagination.svelte';
	import FeedbackSearchForm from './_components/FeedbackSearchForm.svelte';

	export let items: Feedback[] = [];
	export let pages: number = 1;
	export let currentPage: number = 1;
	export let selectedItem: Feedback | null = null;

	let containerRef: HTMLDivElement | null = null;

	function resetScroll() {
		if (containerRef) {
			containerRef.scrollTop = 0;
		}
	}
</script>

<FeedbackSearchForm />

<section class="content">
	<div class="content-scroll" bind:this={containerRef}>
		<div class="h-0">
			{#each items as feedback}
				{#if feedback.type === 'text'}
					<FeedbackListItemText {feedback} active={feedback.id === selectedItem?.id} on:select />
				{:else if feedback.type === 'cnps'}
					<FeedbackListItemCNPS {feedback} active={feedback.id === selectedItem?.id} on:select />
				{:else if feedback.type === 'csat'}
					<FeedbackListItemCSAT {feedback} active={feedback.id === selectedItem?.id} on:select />
				{/if}
			{/each}
		</div>
	</div>
</section>

{#if pages > 1}
	<FeedbackListPagination {pages} {currentPage} on:change={resetScroll} />
{/if}

<style>
	.content {
		display: flex;
		flex: 1;
		min-height: 0px;
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
