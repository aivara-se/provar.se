<script>
	import { page } from '$app/stores';
	import { signIn, signOut } from '@auth/sveltekit/client';
	import { Button } from 'flowbite-svelte';
	import { GithubSolid } from 'flowbite-svelte-icons';

	const signInWithGithub = () => {
		signIn('github', { callbackUrl: '/' });
	};
</script>

<div class="wrapper">
	<main>
		{#if $page.data.session}
			<Button size="sm" color="red" outline on:click={() => signOut()}>Logout</Button>
		{:else}
			<Button size="sm" color="dark" outline on:click={signInWithGithub}>
				<GithubSolid class="w-3 h-3 mr-1" />
				Login with Github
			</Button>
		{/if}
	</main>
</div>

<style>
	.wrapper {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		height: 100vh;
	}

	main {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		text-align: center;
	}
</style>
