<script lang="ts">
	import { ChevronRightIcon } from 'lucide-svelte';
	import type { Feedback } from '$lib/types';

	export let feedbacks: Feedback[] = [];
</script>

{#if feedbacks.length}
	<table class="table">
		<thead>
			<tr>
				<th>Text</th>
				<th class="hidden md:table-cell">Created At</th>
				<th></th>
			</tr>
		</thead>
		<tbody>
			{#each feedbacks as feedback}
				<tr>
					<td>{feedback.data.text}</td>
					<td class="hidden md:table-cell">{feedback.createdAt}</td>
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
