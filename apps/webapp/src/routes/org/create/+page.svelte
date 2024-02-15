<script lang="ts">
	import { AuthLayout } from '$lib/client/layout';
	import route from './route.meta';
	import CreateOrganizationForm from './_components/CreateOrganization.form.svelte';
</script>

<AuthLayout {route}>
	<div class="flex flex-col w-80 gap-4">
		<CreateOrganizationForm />
	</div>
</AuthLayout>
