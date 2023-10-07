<script>
	import { CenteredLayout } from '$lib/ui';
	import { signIn } from '@auth/sveltekit/client';
	import { Button } from 'flowbite-svelte';
	import { GithubSolid, GoogleSolid } from 'flowbite-svelte-icons';
</script>

<CenteredLayout>
	<div class="">
		<Button size="sm" color="red" outline on:click={() => signIn('google', { callbackUrl: '/' })}>
			<GoogleSolid class="w-3 h-3 mr-1" />
			Login with Google
		</Button>
	</div>

	<div class="mt-6">
		<Button size="sm" color="dark" outline on:click={() => signIn('github', { callbackUrl: '/' })}>
			<GithubSolid class="w-3 h-3 mr-1" />
			Login with Github
		</Button>
	</div>
</CenteredLayout>
