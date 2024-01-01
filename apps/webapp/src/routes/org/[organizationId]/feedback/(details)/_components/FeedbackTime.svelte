<script lang="ts">
	import type { Feedback } from '$lib/types';
	import { P } from 'flowbite-svelte';

	export let feedback: Feedback;
</script>

<P size="xs" class="mt-2 text-gray-500">
	Created: {feedback.createdAt.toLocaleDateString()} at {feedback.createdAt.toLocaleTimeString()}
</P>
