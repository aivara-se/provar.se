<script lang="ts">
	import type { CSATFeedback } from '$lib/types';
	import { Heading, Heart, Rating } from 'flowbite-svelte';

	export let feedback: CSATFeedback;

	$: rating = feedback.data.csat * 10;
</script>

<section class="mt-6">
	<Heading customSize="text-md font-semibold">Customer Satisfaction</Heading>
	<div class="mt-3">
		<Rating total={10} {rating} icon={Heart} size={40} />
	</div>
</section>
