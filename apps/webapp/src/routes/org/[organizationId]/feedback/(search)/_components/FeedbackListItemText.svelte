<script lang="ts">
	import type { TextFeedback } from '$lib/types';
	import { Badge, Card, P } from 'flowbite-svelte';
	import { createEventDispatcher } from 'svelte';

	export let feedback: TextFeedback;
	export let active: boolean = false;

	const dispatch = createEventDispatcher();
</script>

<Card
	shadow={false}
	padding="none"
	class="mb-3 py-2 px-3 cursor-pointer hover:border-gray-400 {active ? 'border-gray-400' : ''}"
	on:click={() => dispatch('select', feedback)}
>
	<div class="flex justify-end pointer-events-none">
		<Badge color="none">Text</Badge>
	</div>
	<P size="sm" class="antialiased">“{feedback.data.text}”</P>
	<P size="xs" class="mt-2 text-gray-500">
		{feedback.createdAt.toLocaleDateString()} at {feedback.createdAt.toLocaleTimeString()}
	</P>
</Card>
