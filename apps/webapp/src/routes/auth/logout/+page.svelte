<script lang="ts">
	import { AuthLayout } from '$lib/client/layout';
	import { signOut } from '@auth/sveltekit/client';
</script>

<AuthLayout>
	<div class="flex flex-col w-80">
		<button class="btn btn-md btn-block btn-neutral" on:click={() => signOut()}>Logout</button>
	</div>
</AuthLayout>
