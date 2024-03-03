<script lang="ts">
	import type { Feedback } from '$lib/types';

	export let feedback: Feedback;
</script>

<section class="flex items-center justify-between">
	<p class="text-xs mt-2 text-gray-500 font-semibold">
		Created: {feedback.createdAt.toLocaleDateString()} at {feedback.createdAt.toLocaleTimeString()}
	</p>
</section>
