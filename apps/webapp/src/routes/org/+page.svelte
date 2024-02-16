<script lang="ts">
	import { page } from '$app/stores';
	import { AuthLayout } from '$lib/client/layout';
	import route from './route.meta';

	$: organizations = $page.data.organizations || [];
</script>

<AuthLayout {route}>
	<div class="flex flex-col w-80 gap-4">
		<h1 class="text-lg font-medium antialiased">Organizations</h1>
		{#each organizations as organization}
			<a href="/org/{organization.id}" class="btn btn-block btn-neutral rounded-lg">
				{organization.name}
			</a>
		{/each}
		<div class="divider">or</div>
		<a href="/org/create" class="btn btn-block btn-ghost">Create an organization</a>
	</div>
</AuthLayout>
