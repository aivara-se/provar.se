<script lang="ts">
	import type { Feedback } from '$lib/client/types';
	import { QuoteIcon } from 'lucide-svelte';
	import FeedbackAvatar from './FeedbackAvatar.svelte';

	export let feedback: Feedback;
</script>

<section class="p-4 pr-8 rounded-lg bg-black/10 relative">
	{#if feedback.feedbackData['response-text']}
		<p>“{feedback.feedbackData['response-text']}”</p>
	{:else}
		<p class="italic opacity-50">Feedback text not available</p>
	{/if}

	<FeedbackAvatar user={feedback.feedbackUser} />
	<QuoteIcon class="absolute top-4 right-4 w-4 h-4 opacity-50" />
</section>
