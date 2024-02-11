<script lang="ts">
	import { signOut } from '@auth/sveltekit/client';
	import { AuthLayout } from '$lib/client/layout';
	import route from './route.meta';
</script>

<AuthLayout {route}>
	<div class="flex flex-col w-80">
		<button class="btn btn-block" on:click={() => signOut()}>Logout</button>
	</div>
</AuthLayout>
