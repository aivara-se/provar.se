<script lang="ts">
	import { AuthLayout } from '$lib/client/layout';
	import CreateOrganizationForm from './(forms)/CreateOrganizationForm.svelte';
	import route from './route.meta';

	export let data;
</script>

<AuthLayout {route}>
	<div class="flex flex-col w-80 gap-4">
		<CreateOrganizationForm formData={data.CreateOrganizationForm} />
	</div>
</AuthLayout>
