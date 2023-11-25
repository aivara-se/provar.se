<script>
	import { page } from '$app/stores';
	import { Alert, Breadcrumb, BreadcrumbItem, Heading } from 'flowbite-svelte';
	import { InfoCircleSolid } from 'flowbite-svelte-icons';
</script>

<Breadcrumb class="mb-6">
	<BreadcrumbItem href={`/org/${$page.params.organizationId}`} home>Home</BreadcrumbItem>
	<BreadcrumbItem>Overview</BreadcrumbItem>
</Breadcrumb>

<Heading customSize="mb-2 text-xl font-semibold">Overview</Heading>

<Alert color="blue" rounded={false} class="border-t-4">
	<InfoCircleSolid slot="icon" class="w-4 h-4" />
	This page is under construction.
</Alert>
