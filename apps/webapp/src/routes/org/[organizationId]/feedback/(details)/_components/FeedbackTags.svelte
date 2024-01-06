<script lang="ts">
	import { mergeSearchValue } from '$lib/client/search';
	import type { Feedback } from '$lib/types';
	import { Badge } from 'flowbite-svelte';

	export let feedback: Feedback;

	function filterByType() {
		mergeSearchValue({ type: [feedback.type] });
	}

	function filterByTags(key: string, val: string) {
		mergeSearchValue({ tags: { [key]: val } });
	}
</script>

<section class="mt-3 flex gap-2">
	<a href="#filter-tags" role="button" on:click={filterByType}>
		<Badge color="dark">
			type: {feedback.type === 'cnps' ? 'nps' : feedback.type === 'csat' ? 'csat' : 'text'}
		</Badge>
	</a>

	{#each Object.keys(feedback.tags || {}).sort() as key}
		<a
			href="#filter-tags"
			role="button"
			on:click={() => filterByTags(key, feedback.tags?.[key] || '')}
		>
			<Badge color="blue">
				{key}: {feedback.tags?.[key]}
			</Badge>
		</a>
	{/each}
</section>
