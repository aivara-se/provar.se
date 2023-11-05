<script>
	import { Tabs, TabItem } from 'flowbite-svelte';
	import OrganizationTabContent from './_tab-content/OrganizationTabContent.svelte';
	import MembersTabContent from './_tab-content/MembersTabContent.svelte';
	import CredentialsTabContent from './_tab-content/CredentialsTabContent.svelte';
</script>

<Tabs style="underline">
	<TabItem open title="Organization"><OrganizationTabContent /></TabItem>
	<TabItem title="Members"><MembersTabContent /></TabItem>
	<TabItem title="Credentials"><CredentialsTabContent /></TabItem>
</Tabs>
