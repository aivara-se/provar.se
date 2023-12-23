<script lang="ts">
	import type { CSATFeedback } from '$lib/types';
	import { Badge, Card, P } from 'flowbite-svelte';
	import { ThumbsDownOutline, ThumbsUpOutline } from 'flowbite-svelte-icons';
	import { createEventDispatcher } from 'svelte';

	export let feedback: CSATFeedback;
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
		<Badge color="none">CSAT</Badge>
	</div>
	{#if feedback.data.text}
		<P size="sm" class="antialiased">“{feedback.data.text}”</P>
	{:else}
		<P size="sm" italic class="text-gray-500">No feedback text</P>
	{/if}
	<div class="flex items-center justify-between">
		<P size="xs" class="mt-2 text-gray-500">
			{feedback.createdAt.toLocaleDateString()} at {feedback.createdAt.toLocaleTimeString()}
		</P>
		{#if feedback.data.csat >= 0.9}
			<ThumbsUpOutline class="w-3 h-3 text-green-500 pointer-events-none" />
		{/if}
		{#if feedback.data.csat < 0.5}
			<ThumbsDownOutline class="w-3 h-3 text-rose-500 pointer-events-none" />
		{/if}
	</div>
</Card>
