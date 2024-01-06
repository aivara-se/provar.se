<script lang="ts">
	import type { Feedback } from '$lib/types';
	import { P } from 'flowbite-svelte';

	export let feedback: Feedback;
</script>

<section class="flex items-center justify-between">
	<P size="xs" class="mt-2 text-gray-500 font-semibold">
		Created: {feedback.createdAt.toLocaleDateString()} at {feedback.createdAt.toLocaleTimeString()}
	</P>
</section>
