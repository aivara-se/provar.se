<script lang="ts">
	import type { TextFeedback } from '$lib/types';
	import { format } from 'date-fns';

	export let feedback: TextFeedback;
</script>

<span class="text-xs font-semibold absolute top-2 right-4 opacity-50">Text</span>
<p>{feedback.data.text}</p>
<p class="mt-1 text-xs opacity-50">
	{format(feedback.createdAt, 'yyyy-MM-dd')} at {format(feedback.createdAt, 'HH:mm')}
</p>
