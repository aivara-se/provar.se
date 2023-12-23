<script lang="ts">
	import type { Feedback } from '$lib/types';
	import { Button, ButtonGroup, Input } from 'flowbite-svelte';
	import { SearchOutline } from 'flowbite-svelte-icons';
	import FeedbackListItemCNPS from './FeedbackListItemCNPS.svelte';
	import FeedbackListItemCSAT from './FeedbackListItemCSAT.svelte';
	import FeedbackListItemText from './FeedbackListItemText.svelte';
	import FeedbackListPagination from './FeedbackListPagination.svelte';

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

<section class="filters">
	<div>
		<ButtonGroup class="w-full">
			<Input id="search" size="sm" placeholder="Search ..." />
			<Button color="light">
				<SearchOutline class="w-3 h-3" />
			</Button>
		</ButtonGroup>
	</div>
</section>

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

<FeedbackListPagination {pages} {currentPage} on:change={resetScroll} />

<style>
	.filters {
		display: flex;
		flex-direction: column;
	}

	.content {
		display: flex;
		flex: 1;
		min-height: 0px;
		padding: 20px 0 0 0;
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
