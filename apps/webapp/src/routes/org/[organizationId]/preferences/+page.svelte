<script>
	import { page } from '$app/stores';
	import { Breadcrumb, BreadcrumbItem, Heading, TabItem, Tabs } from 'flowbite-svelte';
	import CredentialsTabContent from './_tab-content/CredentialsTabContent.svelte';
	import MembersTabContent from './_tab-content/MembersTabContent.svelte';
	import OrganizationTabContent from './_tab-content/OrganizationTabContent.svelte';
</script>

<Breadcrumb class="mb-6">
	<BreadcrumbItem href={`/org/${$page.params.organizationId}`} home>Home</BreadcrumbItem>
	<BreadcrumbItem>Getting Started</BreadcrumbItem>
</Breadcrumb>

<Heading customSize="mb-2 text-xl font-semibold">Preferences</Heading>

<Tabs style="underline">
	<TabItem open title="Settings"><OrganizationTabContent /></TabItem>
	<TabItem title="Members"><MembersTabContent /></TabItem>
	<TabItem title="API Keys"><CredentialsTabContent /></TabItem>
</Tabs>
