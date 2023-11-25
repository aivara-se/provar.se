<script>
	import { page } from '$app/stores';
	import { Alert, Breadcrumb, BreadcrumbItem, Button, Heading } from 'flowbite-svelte';
	import { InfoCircleSolid } from 'flowbite-svelte-icons';
</script>

<Breadcrumb class="mb-6">
	<BreadcrumbItem href={`/org/${$page.params.organizationId}`} home>Home</BreadcrumbItem>
	<BreadcrumbItem>Feedbacks</BreadcrumbItem>
</Breadcrumb>

<Heading customSize="mb-2 text-xl font-semibold">Feedbacks</Heading>

<Alert color="blue" rounded={false} class="border-t-4 mb-6">
	<InfoCircleSolid slot="icon" class="w-4 h-4" />
	This page is under construction.
</Alert>

<div>
	<Button
		href={`/org/${$page.params.organizationId}/feedback/project`}
		color="light"
		size="sm"
		outline
	>
		Go to projects
	</Button>
</div>
