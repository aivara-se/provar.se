<script lang="ts">
	import { FeedbackType, type Feedback } from '$lib/types';
	import { format } from 'date-fns';
	import { ChevronRightIcon } from 'lucide-svelte';

	export let feedbacks = {
		count: 0,
		pages: 0,
		items: [] as Feedback[]
	};
</script>

{#if feedbacks.items.length}
	<table class="table mt-4">
		<thead>
			<tr>
				<th class="hidden md:table-cell">Created At</th>
				<th>Type</th>
				<th>Text</th>
				<th></th>
			</tr>
		</thead>
		<tbody>
			{#each feedbacks.items as feedback}
				<tr>
					<td class="hidden md:table-cell">{format(feedback.createdAt, 'yyyy-MM-dd')}</td>
					<td>{feedback.type}</td>
					<td>{feedback.data.text}</td>
					<td class="text-right">
						<a href="./explore/{feedback.id}" role="button" class="btn btn-sm btn-ghost">
							<ChevronRightIcon class="w-4 h-4" />
						</a>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
{/if}
