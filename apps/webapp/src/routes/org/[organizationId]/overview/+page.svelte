<script>
	import { page } from '$app/stores';
	import { Breadcrumb, BreadcrumbItem } from 'flowbite-svelte';
</script>

<section class="flex items-start h-10 justify-between mb-1">
	<section>
		<Breadcrumb class="mb-6">
			<BreadcrumbItem href={`/org/${$page.params.organizationId}`} home>Home</BreadcrumbItem>
			<BreadcrumbItem>Overview</BreadcrumbItem>
		</Breadcrumb>
	</section>
	<section class="flex gap-2 items-start"></section>
</section>
