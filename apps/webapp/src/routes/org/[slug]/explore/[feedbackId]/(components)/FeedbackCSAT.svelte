<script lang="ts">
	import type { Feedback } from '$lib/client/types';
	import { HeartIcon } from 'lucide-svelte';

	interface Details {
		title: string;
		rating: number;
		description: string;
	}

	export let feedback: Feedback;

	function getDetails(data: Record<string, string>): Details {
		const value = parseInt(data['response-data']);
		if (feedback.feedbackData['response-type'] === 'low') {
			return {
				title: 'Low Satisfaction',
				rating: value,
				description:
					'Customers have expressed low satisfaction levels with your product or service. Improvement is needed.'
			};
		} else if (feedback.feedbackData['response-type'] === 'mid') {
			return {
				title: 'Moderate Satisfaction',
				rating: value,
				description:
					'Customers are moderately satisfied, but there is room for enhancement to increase satisfaction levels.'
			};
		} else if (feedback.feedbackData['response-type'] === 'high') {
			return {
				title: 'High Satisfaction',
				rating: value,
				description:
					'Customers are highly satisfied with your product or service. Keep up the good work!'
			};
		}
		return {
			title: 'Invalid Data',
			rating: value,
			description:
				'The Customer Satisfaction Score ranges from 0 to 10. Please provide a valid score within this range.'
		};
	}

	$: details = getDetails(feedback.feedbackData);
</script>

<section class="mt-4">
	<h3 class="text-lg font-semibold">Customer Satisfaction</h3>
	<div class="mt-3">
		<div class="flex">
			{#each [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] as i}
				<HeartIcon
					class="w-6 h-6 text-red-500 {i < details.rating ? 'opacity-100' : 'opacity-20'}"
				/>
			{/each}
		</div>

		<p class="text-sm mt-2">
			<span class="font-semibold">{details.title}:</span>
			<span>{details.description}</span>
		</p>
	</div>
</section>
