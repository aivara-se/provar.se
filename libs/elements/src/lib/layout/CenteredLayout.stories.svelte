<script module>
	import { CenteredLayout, TextInput } from '$lib/index.js';
	import { defineMeta } from '@storybook/addon-svelte-csf';

	const { Story } = defineMeta({
		title: 'Layout/CenteredLayout',
		component: CenteredLayout,
	});
</script>

<Story name="without children" />

<Story name="with some text">
	Hello World
</Story>

<Story name="eg: login page">
	<div class="card w-96 bg-base-100 shadow-xl">
		<div class="card-body">
			<h2 class="card-title text-center justify-center mb-6">Welcome Back</h2>
			<form class="space-y-4">
				<TextInput
					name="email"
					label="Email"
					value=""
					className="input-bordered w-full"
					props={{ type: 'email', placeholder: 'Enter your email', required: true }}
				/>
				<TextInput
					name="password"
					label="Password"
					value=""
					className="input-bordered w-full"
					props={{ type: 'password', placeholder: 'Enter your password', required: true }}
				/>
				<div class="card-actions justify-center pt-4">
					<button class="btn btn-primary w-full">Sign In</button>
				</div>
			</form>
			<div class="divider">OR</div>
			<div class="text-center">
				<a href="#" class="link link-primary text-sm">Forgot your password?</a>
			</div>
			<div class="text-center mt-4">
				<span class="text-sm">Don't have an account? </span>
				<a href="#" class="link link-primary text-sm">Sign up</a>
			</div>
		</div>
	</div>
</Story>
