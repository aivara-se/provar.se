<script lang="ts">
	import type { CSATFeedback } from '$lib/types';
	import { Heading, P } from 'flowbite-svelte';
	import { Rating, Heart } from 'flowbite-svelte';

	interface Details {
		title: string;
		rating: number;
		description: string;
	}

	export let feedback: CSATFeedback;

	function getDetails(value: number): Details {
		if (value >= 0 && value <= 0.6) {
			return {
				title: 'Low Satisfaction',
				rating: value * 10,
				description:
					'Customers have expressed low satisfaction levels with your product or service. Improvement is needed.'
			};
		} else if (value > 0.6 && value <= 0.8) {
			return {
				title: 'Moderate Satisfaction',
				rating: value * 10,
				description:
					'Customers are moderately satisfied, but there is room for enhancement to increase satisfaction levels.'
			};
		} else if (value > 0.8 && value <= 1) {
			return {
				title: 'High Satisfaction',
				rating: value * 10,
				description:
					'Customers are highly satisfied with your product or service. Keep up the good work!'
			};
		}
		return {
			title: 'Invalid Range',
			rating: value * 10,
			description:
				'The Customer Satisfaction Score ranges from 0 to 1. Please provide a valid score within this range.'
		};
	}

	$: details = getDetails(feedback.data.csat);
</script>

<section class="mt-6">
	<Heading customSize="text-md font-semibold">Customer Satisfaction</Heading>
	<div class="mt-3">
		<Rating total={10} rating={details.rating} icon={Heart} size={30} />
		<P class="text-sm mt-2">
			<span class="text-gray-700 font-semibold">{details.title}:</span>
			<span class="text-gray-600">{details.description}</span>
		</P>
	</div>
</section>
