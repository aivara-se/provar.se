<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import CredentialsTable from './(components)/CredentialsTable.svelte';
	import CreateCredentialForm from './(forms)/CreateCredentialForm.svelte';
	import route from './route.meta';

	export let data;
</script>

<DashLayout {route}>
	<CredentialsTable credentials={data.credentials} />
	<CreateCredentialForm formData={data.CreateCredentialForm} />
</DashLayout>
