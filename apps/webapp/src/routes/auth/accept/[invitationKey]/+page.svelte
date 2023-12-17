<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { CenteredLayout } from '$lib/client/ui';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button } from 'flowbite-svelte';

	async function acceptInvitation() {
		const data = new FormData();
		const response = await fetch('', {
			method: 'post',
			body: data
		});
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
	}
</script>

<CenteredLayout>
	<Button size="sm" color="red" outline on:click={acceptInvitation}>Accept Invitation</Button>
</CenteredLayout>
