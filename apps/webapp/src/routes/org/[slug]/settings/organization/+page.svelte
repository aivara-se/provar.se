<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import DeleteOrganizationBlock from './(components)/DeleteOrganizationBlock.svelte';
	import LeaveOrganizationBlock from './(components)/LeaveOrganizationBlock.svelte';
	import UpdateOrganizationForm from './(forms)/UpdateOrganizationForm.svelte';
	import route from './route.meta';

	export let data;

	// We know this is not null
	$: currentUser = data.user!;
</script>

<DashLayout {route}>
	<UpdateOrganizationForm organization={data.organization} formData={data.UpdateOrganizationForm} />
	<LeaveOrganizationBlock user={currentUser} organization={data.organization} />
	<DeleteOrganizationBlock organization={data.organization} />
</DashLayout>
