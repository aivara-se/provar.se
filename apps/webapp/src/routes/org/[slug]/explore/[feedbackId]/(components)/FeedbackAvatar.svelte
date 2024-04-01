<script lang="ts">
	import type { Feedback } from '$lib/client/types';

	export let user: Feedback['feedbackUser'];

	$: name = user?.name || 'Anonymous';
</script>

<div class="mt-3 flex items-center gap-2">
	{#if user?.avatar}
		<div class="avatar">
			<div class="w-6 h-6 rounded-full">
				<img src={user?.avatar} alt={user?.name} />
			</div>
		</div>
	{:else}
		<div class="avatar placeholder">
			<div class="w-6 h-6 rounded-full">
				<span class="text-xs">{name.slice(2)}</span>
			</div>
		</div>
	{/if}
	<cite class="text-sm font-semibold">{name}</cite>
</div>
