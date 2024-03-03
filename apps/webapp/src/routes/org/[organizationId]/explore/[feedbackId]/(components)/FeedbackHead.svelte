<script lang="ts">
	import type { Feedback } from '$lib/types';
	import { format } from 'date-fns';

	export let feedback: Feedback;
</script>

<section class="flex items-center justify-between">
	<p class="text-xs mt-2 text-gray-500 font-semibold">
		Created: {format(feedback.createdAt, 'yyyy-MM-dd')} at {format(feedback.createdAt, 'HH:mm')}
	</p>
</section>
