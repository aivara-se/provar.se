<script>
	import { page } from '$app/stores';
	import { Breadcrumb, BreadcrumbItem, TabItem, Tabs } from 'flowbite-svelte';
	import CredentialsTabContent from './(credentials)/CredentialsTabContent.svelte';
	import MembersTabContent from './(members)/MembersTabContent.svelte';
	import OrganizationTabContent from './(organization)/OrganizationTabContent.svelte';
</script>

<section class="flex items-start h-10 justify-between mb-1">
	<section>
		<Breadcrumb class="mb-6">
			<BreadcrumbItem href={`/org/${$page.params.organizationId}`} home>Home</BreadcrumbItem>
			<BreadcrumbItem>Preferences</BreadcrumbItem>
		</Breadcrumb>
	</section>
	<section class="flex gap-2 items-start"></section>
</section>

<Tabs style="underline">
	<TabItem open title="Organization"><OrganizationTabContent /></TabItem>
	<TabItem title="Members"><MembersTabContent /></TabItem>
	<TabItem title="API Keys"><CredentialsTabContent /></TabItem>
</Tabs>
