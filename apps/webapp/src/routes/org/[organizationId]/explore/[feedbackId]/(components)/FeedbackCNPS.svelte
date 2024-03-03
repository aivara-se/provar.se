<script lang="ts">
	import type { CNPSFeedback } from '$lib/types';
	import { StarIcon } from 'lucide-svelte';

	interface Details {
		title: string;
		rating: number;
		description: string;
	}

	export let feedback: CNPSFeedback;

	function getDetails(value: number): Details {
		if (value >= 0 && value <= 0.6) {
			return {
				title: 'Detractor',
				rating: Math.round(value * 10),
				description:
					'Detractors are unhappy customers who can damage your brand and impede growth. They are unlikely to recommend your product or service.'
			};
		} else if (value > 0.6 && value <= 0.8) {
			return {
				title: 'Passive',
				rating: Math.round(value * 10),
				description:
					'Passives are satisfied but unenthusiastic customers. They may switch to competitors and are less likely to spread positive word-of-mouth.'
			};
		} else if (value > 0.8 && value <= 1) {
			return {
				title: 'Promoter',
				rating: Math.round(value * 10),
				description:
					'Promoters are your loyal and enthusiastic customers. They love your product or service and are likely to recommend it to others.'
			};
		}
		return {
			title: 'Invalid Range',
			rating: Math.round(value * 10),
			description:
				'The Net Promoter Score ranges from 0 to 1. Please provide a valid score within this range.'
		};
	}

	$: details = getDetails(feedback.data.cnps);
</script>

<section class="mt-4">
	<h3 class="text-lg font-semibold">Net Promoter Score</h3>
	<div class="mt-3">
		<div class="flex">
			{#each [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] as i}
				<StarIcon
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
