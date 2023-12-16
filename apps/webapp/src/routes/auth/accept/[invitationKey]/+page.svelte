<script>
	import { CenteredLayout } from '$lib/ui';
	import { Button } from 'flowbite-svelte';

	async function acceptInvitation() {
		console.log('acceptInvitation');
	}
</script>

<CenteredLayout>
	<Button size="sm" color="red" outline on:click={() => acceptInvitation()}>
		Accept Invitation
	</Button>
</CenteredLayout>
