<script>
	import { CenteredLayout } from '$lib/ui';
	import { signIn } from '@auth/sveltekit/client';
	import { Button } from 'flowbite-svelte';
	import { GithubSolid } from 'flowbite-svelte-icons';
</script>

<CenteredLayout>
	<Button size="sm" color="dark" outline on:click={() => signIn('github', { callbackUrl: '/' })}>
		<GithubSolid class="w-3 h-3 mr-1" />
		Login with Github
	</Button>
</CenteredLayout>
