<script>
	import { CenteredLayout } from '$lib/ui';
	import { signOut } from '@auth/sveltekit/client';
	import { Button } from 'flowbite-svelte';
</script>

<CenteredLayout>
	<Button size="sm" color="red" outline on:click={() => signOut()}>Logout</Button>
</CenteredLayout>
