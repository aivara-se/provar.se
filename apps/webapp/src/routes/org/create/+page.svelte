<script lang="ts">
	import { AuthLayout } from '$lib/client/layout';
	import CreateOrganizationForm from './(forms)/CreateOrganizationForm.svelte';

	export let data;
</script>

<AuthLayout>
	<div class="flex flex-col w-80 gap-4">
		<CreateOrganizationForm formData={data.CreateOrganizationForm} />
	</div>
</AuthLayout>
