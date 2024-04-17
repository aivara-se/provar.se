<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import FeedbackCNPS from './(components)/FeedbackCNPS.svelte';
	import FeedbackCSAT from './(components)/FeedbackCSAT.svelte';
	import FeedbackHead from './(components)/FeedbackHead.svelte';
	import FeedbackMeta from './(components)/FeedbackMeta.svelte';
	import FeedbackTags from './(components)/FeedbackTags.svelte';
	import FeedbackText from './(components)/FeedbackText.svelte';
	import route from './route.meta';

	export let data;
</script>

<DashLayout {route} user={data.user} organizations={data.organizations}>
	<section>
		<FeedbackText feedback={data.feedback} />
		<FeedbackHead feedback={data.feedback} />
		<FeedbackTags organization={data.organization} feedback={data.feedback} />

		{#if data.feedback.feedbackType === 'cnps'}
			<FeedbackCNPS feedback={data.feedback} />
		{/if}

		{#if data.feedback.feedbackType === 'csat'}
			<FeedbackCSAT feedback={data.feedback} />
		{/if}

		{#if data.feedback.feedbackMeta && Object.keys(data.feedback.feedbackMeta).length}
			<FeedbackMeta organization={data.organization} feedback={data.feedback} />
		{/if}
	</section>
</DashLayout>
