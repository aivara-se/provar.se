<script lang="ts">
	import { page } from '$app/stores';
	import { verifyEmailLogin } from '$lib/client/auth';
	import { AuthLayout } from '$lib/client/layout';
	import { onMount } from 'svelte';

	onMount(() => {
		verifyEmailLogin($page.params.token);
	});
</script>

<AuthLayout>
	<div class="flex flex-col prose text-center">
		<h2>Verifying email login</h2>
		<p>You will be redirected to the application soon.</p>
	</div>
</AuthLayout>
