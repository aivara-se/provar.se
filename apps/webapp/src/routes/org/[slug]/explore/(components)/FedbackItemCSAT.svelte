<script lang="ts">
	import type { CSATFeedback } from '$lib/client/types';
	import { format } from 'date-fns';
	import { ThumbsDownIcon, ThumbsUpIcon } from 'lucide-svelte';

	export let feedback: CSATFeedback;
</script>

<span class="text-xs font-semibold absolute top-2 right-4 opacity-50">CSAT</span>
<p>{feedback.data.text}</p>
<p class="mt-1 text-xs opacity-50">
	{format(feedback.createdAt, 'yyyy-MM-dd')} at {format(feedback.createdAt, 'HH:mm')}
</p>
<div class="absolute bottom-3 right-4">
	{#if feedback.data.csat >= 0.9}
		<ThumbsUpIcon class="w-3 h-3" />
	{/if}
	{#if feedback.data.csat < 0.5}
		<ThumbsDownIcon class="w-3 h-3" />
	{/if}
</div>
