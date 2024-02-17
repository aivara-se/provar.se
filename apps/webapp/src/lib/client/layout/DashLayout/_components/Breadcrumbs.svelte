<script lang="ts">
	import { page } from '$app/stores';
	import { getBreadcrumbs, getHomeUrl, type Route } from '$lib/client/routes';
	import { HomeIcon } from 'lucide-svelte';

	export let route: Route;

	$: homeLink = getHomeUrl($page.params);
	$: segments = getBreadcrumbs(route).map((r) => ({
		id: r.id,
		name: r.getName($page.params),
		path: r.getPath($page.params)
	}));
</script>

<div class="text-sm antialiased breadcrumbs">
	<ul>
		<li>
			<a href={homeLink}><HomeIcon class="w-4 h-4" /></a>
		</li>
		{#each segments as segment (segment.id)}
			<li>
				<a href={segment.path}>{segment.name}</a>
			</li>
		{/each}
	</ul>
</div>
