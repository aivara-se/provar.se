<script lang="ts">
	import type { Feedback } from '$lib/types';
	import { Badge } from 'flowbite-svelte';
	import { TagOutline } from 'flowbite-svelte-icons';

	export let feedback: Feedback;
</script>

<section class="mt-3 flex gap-2">
	{#each Object.keys(feedback.tags || {}).sort() as key}
		<Badge border color="dark" class="border-gray-300 text-gray-600">
			<TagOutline class="w-2.5 h-2.5 me-1.5" />
			{key}: {feedback.tags?.[key]}
		</Badge>
	{/each}
</section>
