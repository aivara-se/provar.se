<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import InvitationsTable from './(components)/InvitationsTable.svelte';
	import MembersTable from './(components)/MembersTable.svelte';
	import InviteMemberForm from './(forms)/InviteMemberForm.svelte';
	import route from './route.meta';

	export let data;
</script>

<DashLayout {route} user={data.user} organizations={data.organizations}>
	<MembersTable members={data.members} />
	<InvitationsTable invitations={data.invitations} />
	<InviteMemberForm organization={data.organization} formData={data.InviteMemberForm} />
</DashLayout>
