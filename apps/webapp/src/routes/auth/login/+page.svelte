<script lang="ts">
	import { signIn } from '@auth/sveltekit/client';
	import { AuthLayout } from '$lib/client/layout';
	import route from './route.meta';
	import GoogleIcon from './GoogleIcon.svelte';
	import GithubIcon from './GithubIcon.svelte';

	let email: string;
</script>

<AuthLayout {route}>
	<div class="flex flex-col w-80 gap-4">
		<form
			class="flex flex-col gap-4"
			on:submit={() => signIn('email', { email, callbackUrl: '/' })}
		>
			<label class="form-control w-full">
				<div class="label">
					<span class="label-text">Email</span>
				</div>
				<input type="text" name="email" class="input input-bordered w-full" bind:value={email} />
			</label>
			<input type="submit" class="btn btn-md btn-block btn-neutral" value="Login with email" />
		</form>

		<div class="divider">or</div>

		<button
			class="btn btn-md btn-block btn-neutral"
			on:click={() => signIn('google', { callbackUrl: '/' })}
		>
			<GoogleIcon /> Login with Google
		</button>

		<button
			class="btn btn-md btn-block btn-neutral"
			on:click={() => signIn('github', { callbackUrl: '/' })}
		>
			<GithubIcon /> Login with Github
		</button>
	</div>
</AuthLayout>
