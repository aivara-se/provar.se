<script lang="ts">
	import type { Feedback } from '$lib/client/types';
	import { format } from 'date-fns';

	export let feedback: Feedback;
</script>

<span class="text-xs font-semibold absolute top-2 right-4 opacity-50">CSAT</span>
<p>{feedback.feedbackData['response-text']}</p>
<p class="mt-1 text-xs opacity-50">
	{format(feedback.createdAt, 'yyyy-MM-dd')} at {format(feedback.createdAt, 'HH:mm')}
</p>
