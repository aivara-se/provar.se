<script>
	import { page } from '$app/stores';
	import { Alert, Breadcrumb, BreadcrumbItem, Heading } from 'flowbite-svelte';
	import { InfoCircleSolid } from 'flowbite-svelte-icons';

	$: project = $page.data.project;
</script>

<Breadcrumb class="mb-6">
	<BreadcrumbItem href={`/org/${$page.params.organizationId}`} home>Home</BreadcrumbItem>
	<BreadcrumbItem href={`/org/${$page.params.organizationId}/project`}>Projects</BreadcrumbItem>
	<BreadcrumbItem>{project?.name}</BreadcrumbItem>
</Breadcrumb>

<Heading customSize="mb-2 text-xl font-semibold">{project?.name}</Heading>

<Alert color="blue" rounded={false} class="border-t-4">
	<InfoCircleSolid slot="icon" class="w-4 h-4" />
	This page is under construction.
</Alert>
