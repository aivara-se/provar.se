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

<DashLayout {route}>
	<section>
		<FeedbackHead feedback={data.feedback} />
		<FeedbackText feedback={data.feedback} />
		<FeedbackTags feedback={data.feedback} />

		{#if data.feedback.type === 'cnps'}
			<FeedbackCNPS feedback={data.feedback} />
		{/if}

		{#if data.feedback.type === 'csat'}
			<FeedbackCSAT feedback={data.feedback} />
		{/if}

		{#if data.feedback.meta && Object.keys(data.feedback.meta).length}
			<FeedbackMeta feedback={data.feedback} />
		{/if}
	</section>
</DashLayout>
