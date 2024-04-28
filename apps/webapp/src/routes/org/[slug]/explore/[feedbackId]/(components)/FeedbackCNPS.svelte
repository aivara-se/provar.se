<script lang="ts">
	import type { Feedback } from '$lib/client/types';
	import { StarIcon } from 'lucide-svelte';

	interface Details {
		title: string;
		rating: number;
		description: string;
	}

	export let feedback: Feedback;

	function getDetails(data: Record<string, string>): Details {
		const value = parseInt(data['response-data']);
		if (feedback.feedbackData['response-type'] === 'detractor') {
			return {
				title: 'Detractor',
				rating: value,
				description:
					'Detractors are unhappy customers who can damage your brand and impede growth. They are unlikely to recommend your product or service.'
			};
		} else if (feedback.feedbackData['response-type'] === 'passive') {
			return {
				title: 'Passive',
				rating: value,
				description:
					'Passives are satisfied but unenthusiastic customers. They may switch to competitors and are less likely to spread positive word-of-mouth.'
			};
		} else if (feedback.feedbackData['response-type'] === 'promoter') {
			return {
				title: 'Promoter',
				rating: value,
				description:
					'Promoters are your loyal and enthusiastic customers. They love your product or service and are likely to recommend it to others.'
			};
		}
		return {
			title: 'Invalid Data',
			rating: value,
			description:
				'The Customer Net Promoter Score ranges from 0 to 10. Please provide a valid score within this range.'
		};
	}

	$: details = getDetails(feedback.feedbackData);
</script>

<section class="mt-4">
	<h3 class="text-lg font-semibold">Net Promoter Score</h3>
	<div class="mt-3">
		<div class="flex">
			{#each [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] as i}
				<StarIcon
					class="w-6 h-6 text-yellow-500 {i < details.rating ? 'opacity-100' : 'opacity-20'}"
				/>
			{/each}
		</div>

		<p class="text-sm mt-2">
			<span class="font-semibold">{details.title}:</span>
			<span>{details.description}</span>
		</p>
	</div>
</section>
