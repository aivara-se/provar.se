<script lang="ts">
	import { CenteredLayout, HorizontalSeparator } from '$lib/client/ui';
	import { signIn } from '@auth/sveltekit/client';
	import { Button, Input } from 'flowbite-svelte';
	import { GithubSolid, GoogleSolid } from 'flowbite-svelte-icons';

	let email: string;
</script>

<CenteredLayout>
	<div class="form">
		<Input
			class="rounded-none rounded-l-md"
			id="email"
			bind:value={email}
			placeholder="myname@email.com"
		/>
		<Button
			class="rounded-none rounded-r-md"
			size="sm"
			on:click={() => signIn('email', { email, callbackUrl: '/' })}
		>
			Login
		</Button>
	</div>

	<HorizontalSeparator text="or" />

	<div class="provider">
		<Button size="sm" color="red" outline on:click={() => signIn('google', { callbackUrl: '/' })}>
			<GoogleSolid class="w-3 h-3 mr-1" />
			Login with Google
		</Button>
	</div>

	<div class="provider">
		<Button size="sm" color="dark" outline on:click={() => signIn('github', { callbackUrl: '/' })}>
			<GithubSolid class="w-3 h-3 mr-1" />
			Login with Github
		</Button>
	</div>
</CenteredLayout>

<style>
	.form {
		display: flex;
		width: 250px;
		text-align: left;
	}

	.provider {
		width: 250px;
		display: flex;
		flex-direction: column;
	}

	.provider + .provider {
		margin-top: 20px;
	}
</style>
