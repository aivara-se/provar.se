<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button } from 'flowbite-svelte';

	async function handler() {
		const response = await fetch('', { method: 'post', body: new FormData() });
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
	}
</script>

<section>
	<Button type="submit" size="sm" color="red" on:click={handler}>Accept Invitation</Button>
</section>
