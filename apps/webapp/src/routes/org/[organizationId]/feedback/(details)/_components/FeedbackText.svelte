<script lang="ts">
	import type { Feedback } from '$lib/types';
	import { Blockquote } from 'flowbite-svelte';
	import { QuoteSolid } from 'flowbite-svelte-icons';
	import FeedbackUserAvatar from './FeedbackUserAvatar.svelte';

	export let feedback: Feedback;
</script>

<section class="mt-2">
	<Blockquote bg border baseClass="p-2 px-4">
		<div class="flex justify-end">
			<QuoteSolid class="w-5 h-5 text-gray-400" />
		</div>
		“{feedback.data.text || 'No feedback text'}”
		<FeedbackUserAvatar user={feedback.user} withName={true} />
	</Blockquote>
</section>
