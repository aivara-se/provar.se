<script>
	import { page } from '$app/stores';
	import { signIn, signOut } from '@auth/sveltekit/client';
	import { Button } from 'flowbite-svelte';
	import { GithubSolid } from 'flowbite-svelte-icons';
	import { CenteredLayout } from '$lib/ui';
</script>

<CenteredLayout>
	{#if $page.data.session}
		<Button size="sm" color="red" outline on:click={() => signOut()}>Logout</Button>
	{:else}
		<Button size="sm" color="dark" outline on:click={() => signIn('github', { callbackUrl: '/' })}>
			<GithubSolid class="w-3 h-3 mr-1" />
			Login with Github
		</Button>
	{/if}
</CenteredLayout>
