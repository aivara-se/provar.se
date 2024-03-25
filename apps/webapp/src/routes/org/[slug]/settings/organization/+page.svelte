<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import UpdateOrganizationForm from './(forms)/UpdateOrganizationForm.svelte';
	import route from './route.meta';

	export let data;
</script>

<DashLayout {route}>
	<UpdateOrganizationForm organization={data.organization} formData={data.UpdateOrganizationForm} />
	<!-- <LeaveOrganizationBlock organization={data.organization} /> -->
	<!-- <DeleteOrganizationBlock organization={data.organization} /> -->
</DashLayout>
