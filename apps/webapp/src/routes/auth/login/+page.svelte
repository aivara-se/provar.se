<script lang="ts">
	import { signInWithEmail, signInWithGithub, signInWithGoogle } from '$lib/client/auth';
	import { AuthLayout } from '$lib/client/layout';
	import GithubIcon from './GithubIcon.svelte';
	import GoogleIcon from './GoogleIcon.svelte';

	let email: string;
</script>

<AuthLayout>
	<div class="flex flex-col w-80 gap-4">
		<form class="flex flex-col gap-4" on:submit={() => signInWithEmail(email)}>
			<label class="form-control w-full">
				<div class="label">
					<span class="label-text">Email</span>
				</div>
				<input type="text" name="email" class="input input-bordered w-full" bind:value={email} />
			</label>
			<input type="submit" class="btn btn-md btn-block btn-neutral" value="Login with email" />
		</form>

		<div class="divider">or</div>

		<button class="btn btn-md btn-block btn-neutral" on:click={() => signInWithGoogle()}>
			<GoogleIcon /> Login with Google
		</button>

		<button class="btn btn-md btn-block btn-neutral" on:click={() => signInWithGithub()}>
			<GithubIcon /> Login with Github
		</button>
	</div>
</AuthLayout>
