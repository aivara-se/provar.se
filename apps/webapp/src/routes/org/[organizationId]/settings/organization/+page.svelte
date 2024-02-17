<script lang="ts">
	import { ConfirmModal } from '$lib/client/forms';
	import { DashLayout } from '$lib/client/layout';
	import { AlertCircleIcon } from 'lucide-svelte';
	import UpdateOrganizationForm from './(forms)/UpdateOrganizationForm.svelte';
	import route from './route.meta';
	import LeaveOrganizationBlock from './(components)/LeaveOrganizationBlock.svelte';
	import DeleteOrganizationBlock from './(components)/DeleteOrganizationBlock.svelte';

	export let data;
</script>

<DashLayout {route}>
	<UpdateOrganizationForm formData={data.UpdateOrganizationForm} />
	<LeaveOrganizationBlock organization={data.organization} />
	<DeleteOrganizationBlock organization={data.organization} />
</DashLayout>
