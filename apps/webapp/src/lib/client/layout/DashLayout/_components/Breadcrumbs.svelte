<script lang="ts">
	import { page } from '$app/stores';
	import { getBreadcrumbs, getHomeUrl, type Route } from '$lib/client/routes';
	import { HomeIcon } from 'lucide-svelte';

	export let route: Route;

	$: organizationLink = getHomeUrl($page.params);
	$: organizationName = $page.data?.organization?.name || '';

	$: segments = getBreadcrumbs(route).map((r) => ({
		id: r.id,
		name: r.getName($page.params),
		path: r.getPath($page.params)
	}));
</script>

<div class="text-sm antialiased breadcrumbs items-end">
	<ul>
		<li>
			<a href="/"><HomeIcon class="w-4 h-4" /></a>
		</li>
		<li>
			<a href={organizationLink}>{organizationName}</a>
		</li>
		{#each segments as segment (segment.id)}
			<li>
				<a href={segment.path}>{segment.name}</a>
			</li>
		{/each}
	</ul>
</div>
