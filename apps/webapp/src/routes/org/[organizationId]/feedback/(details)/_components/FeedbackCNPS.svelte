<script lang="ts">
	import type { CNPSFeedback } from '$lib/types';
	import { Heading } from 'flowbite-svelte';
	import { Rating, Heart } from 'flowbite-svelte';

	export let feedback: CNPSFeedback;

	$: rating = feedback.data.cnps * 10;
</script>

<section class="mt-6">
	<Heading customSize="text-md font-semibold">Net Promoter Score</Heading>
	<div class="mt-3">
		<Rating total={10} {rating} icon={Heart} size={40} />
	</div>
</section>
