<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import InvitationsTable from './(components)/InvitationsTable.svelte';
	import MembersTable from './(components)/MembersTable.svelte';
	import InviteMemberForm from './(forms)/InviteMemberForm.svelte';
	import route from './route.meta';

	export let data;
</script>

<DashLayout {route}>
	<MembersTable members={data.members} />
	<InvitationsTable invitations={data.invitations} />
	<InviteMemberForm formData={data.InviteMemberForm} />
</DashLayout>
