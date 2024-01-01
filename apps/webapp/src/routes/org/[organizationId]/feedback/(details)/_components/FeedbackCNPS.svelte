<script lang="ts">
	import type { CNPSFeedback } from '$lib/types';
	import { Heading, P } from 'flowbite-svelte';
	import { Rating, Heart } from 'flowbite-svelte';

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
				rating: value * 10,
				description:
					'Detractors are unhappy customers who can damage your brand and impede growth. They are unlikely to recommend your product or service.'
			};
		} else if (value >= 0.7 && value <= 0.8) {
			return {
				title: 'Passive',
				rating: value * 10,
				description:
					'Passives are satisfied but unenthusiastic customers. They may switch to competitors and are less likely to spread positive word-of-mouth.'
			};
		} else if (value >= 0.9 && value <= 1) {
			return {
				title: 'Promoter',
				rating: value * 10,
				description:
					'Promoters are your loyal and enthusiastic customers. They love your product or service and are likely to recommend it to others.'
			};
		}
		return {
			title: 'Invalid Range',
			rating: value * 10,
			description:
				'The Net Promoter Score ranges from 0 to 1. Please provide a valid score within this range.'
		};
	}

	$: details = getDetails(feedback.data.cnps);
</script>

<section class="mt-6">
	<Heading customSize="text-md font-semibold">Net Promoter Score</Heading>
	<div class="mt-3">
		<Rating total={10} rating={details.rating} icon={Heart} size={30} />
		<P class="text-sm mt-2">
			<span class="text-gray-700 font-semibold">{details.title}:</span>
			<span class="text-gray-600">{details.description}</span>
		</P>
	</div>
</section>
