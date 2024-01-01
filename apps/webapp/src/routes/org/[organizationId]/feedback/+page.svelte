<script lang="ts">
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import type { Feedback } from '$lib/types';
	import { Breadcrumb, BreadcrumbItem, Heading, P } from 'flowbite-svelte';
	import DateRangeSelector from './_components/DateRangeSelector.svelte';
	import FeedbackList from './(search)/FeedbackList.svelte';
	import FeedbackPageActions from './_components/FeedbackPageActions.svelte';
	import FeedbackDetail from './(detail)/FeedbackDetail.svelte';

	$: feedbacks = $page.data.feedbacks;
	$: currentPage = Number.parseInt($page.url.searchParams.get('page') || '1');

	let selectedFeedback: Feedback | null = null;

	$: if (browser) {
		const id = $page.url.searchParams.get('id');
		if (!selectedFeedback && id) {
			const feedback = feedbacks.items.find((f: Feedback) => f.id === id);
			if (feedback) {
				selectFeedback(feedback);
			}
		}
		if (!selectedFeedback && feedbacks.items.length > 0) {
			selectFeedback(feedbacks.items[0]);
		}
	}

	function selectFeedback(feedback: Feedback) {
		selectedFeedback = feedback;
		$page.url.searchParams.set('id', feedback.id);
		goto($page.url.toString());
	}
</script>

<Breadcrumb class="mb-6">
	<BreadcrumbItem href={`/org/${$page.params.organizationId}`} home>Home</BreadcrumbItem>
	<BreadcrumbItem>Feedback</BreadcrumbItem>
</Breadcrumb>

<section class="flex">
	<Heading class="flex-1" customSize="text-xl font-semibold">Feedback</Heading>
	<section class="flex gap-2 items-center">
		<DateRangeSelector />
		<FeedbackPageActions />
	</section>
</section>

<section class="flex flex-1">
	<aside class="sidebar">
		<FeedbackList
			items={feedbacks.items}
			pages={feedbacks.pages}
			{currentPage}
			selectedItem={selectedFeedback}
			on:select={(event) => selectFeedback(event.detail)}
		/>
	</aside>
	<main class="content">
		<FeedbackDetail feedback={selectedFeedback} />
	</main>
</section>

<style>
	.sidebar {
		display: flex;
		flex-shrink: 0;
		padding: 20px 20px 0 0;
		width: 360px;
		flex-direction: column;
		box-sizing: border-box;
		border-right: solid 1px #e0e0e0;
	}

	.content {
		display: flex;
		flex-grow: 1;
	}
</style>
