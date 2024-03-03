<script lang="ts">
	import type { Feedback } from '$lib/types';
	import { QuoteIcon } from 'lucide-svelte';
	import FeedbackAvatar from './FeedbackAvatar.svelte';

	export let feedback: Feedback;
</script>

<section class="mt-2 p-4 pr-8 rounded-lg bg-black/5 dark:bg-black/10 relative">
	<p>
		“{feedback.data.text || 'No feedback text'}”
	</p>

	<FeedbackAvatar user={feedback.user} />
	<QuoteIcon class="absolute top-4 right-4 w-4 h-4 opacity-50" />
</section>
