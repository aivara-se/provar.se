<script>
	import { page } from '$app/stores';
	import { Breadcrumb, BreadcrumbItem, Heading, TabItem, Tabs } from 'flowbite-svelte';
	import CredentialsTabContent from './credentials/CredentialsTabContent.svelte';
	import MembersTabContent from './members/MembersTabContent.svelte';
	import OrganizationTabContent from './organization/OrganizationTabContent.svelte';
</script>

<Breadcrumb class="mb-6">
	<BreadcrumbItem href={`/org/${$page.params.organizationId}`} home>Home</BreadcrumbItem>
	<BreadcrumbItem>Preferences</BreadcrumbItem>
</Breadcrumb>

<Heading customSize="mb-2 text-xl font-semibold">Preferences</Heading>

<Tabs style="underline">
	<TabItem open title="Organization"><OrganizationTabContent /></TabItem>
	<TabItem title="Members"><MembersTabContent /></TabItem>
	<TabItem title="API Keys"><CredentialsTabContent /></TabItem>
</Tabs>
