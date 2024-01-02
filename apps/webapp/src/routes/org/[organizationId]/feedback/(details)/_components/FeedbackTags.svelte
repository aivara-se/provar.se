<script lang="ts">
	import type { Feedback } from '$lib/types';
	import { Badge } from 'flowbite-svelte';
	import { TagOutline } from 'flowbite-svelte-icons';
	import { mergeSearchValue } from '$lib/client/search';

	export let feedback: Feedback;

	function filterByTags(key: string, val: string) {
		mergeSearchValue({ tags: { [key]: val } });
	}
</script>

<section class="mt-3 flex gap-2">
	{#each Object.keys(feedback.tags || {}).sort() as key}
		<a
			href="#filter-tags"
			role="button"
			on:click={() => filterByTags(key, feedback.tags?.[key] || '')}
		>
			<Badge color="blue">
				<TagOutline class="w-2.5 h-2.5 me-1.5" />
				{key}: {feedback.tags?.[key]}
			</Badge>
		</a>
	{/each}
</section>
