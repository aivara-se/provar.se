<script lang="ts">
	import { mergeSearchValue } from '$lib/client/search';
	import type { Feedback } from '$lib/types';
	import { Badge, P } from 'flowbite-svelte';

	export let feedback: Feedback;

	function filterByType() {
		mergeSearchValue({ type: [feedback.type] });
	}
</script>

<section class="flex items-center justify-between">
	<P size="xs" class="mt-2 text-gray-500 font-semibold">
		Created: {feedback.createdAt.toLocaleDateString()} at {feedback.createdAt.toLocaleTimeString()}
	</P>
	<P size="xs" class="mt-2">
		<a href="#filter-type" role="button" on:click={filterByType}>
			<Badge rounded border color="dark">
				{feedback.type === 'cnps' ? 'NPS' : feedback.type === 'csat' ? 'CSAT' : 'Text'}
			</Badge>
		</a>
	</P>
</section>
