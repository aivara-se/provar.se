<script lang="ts">
	import { signOut } from '$lib/client/auth';
	import { AuthLayout } from '$lib/client/layout';
</script>

<AuthLayout>
	<div class="flex flex-col w-80">
		<button class="btn btn-md btn-block" on:click={() => signOut()}>Logout</button>
	</div>
</AuthLayout>
