<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import DeleteOrganizationBlock from './(components)/DeleteOrganizationBlock.svelte';
	import LeaveOrganizationBlock from './(components)/LeaveOrganizationBlock.svelte';
	import UpdateOrganizationForm from './(forms)/UpdateOrganizationForm.svelte';
	import route from './route.meta';

	export let data;
</script>

<DashLayout {route}>
	<UpdateOrganizationForm formData={data.UpdateOrganizationForm} />
	<LeaveOrganizationBlock organization={data.organization} />
	<DeleteOrganizationBlock organization={data.organization} />
</DashLayout>
