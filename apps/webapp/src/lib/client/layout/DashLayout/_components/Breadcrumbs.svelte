<script lang="ts">
	import { page } from '$app/stores';
	import { getBreadcrumbs, getHomeUrl, type Route } from '$lib/client/routes';
	import { HomeIcon, ChevronLeftIcon } from 'lucide-svelte';

	export let route: Route;

	$: homeLink = getHomeUrl($page.params);
	$: segments = getBreadcrumbs(route).map((r) => ({
		id: r.id,
		name: r.getName($page.params),
		path: r.getPath($page.params)
	}));

	function handleBack() {
		history.back();
	}
</script>

<button class="btn btn-sm pl-1 pr-2 gap-0.5 btn-ghost md:hidden" on:click={handleBack}>
	<ChevronLeftIcon class="w-3.5 h-3.5" />Back
</button>

<div class="text-sm antialiased breadcrumbs hidden md:flex">
	<ul>
		<li>
			<a href={homeLink}><HomeIcon class="w-3.5 h-3.5" /></a>
		</li>
		{#each segments as segment (segment.id)}
			<li>
				<a href={segment.path}>{segment.name}</a>
			</li>
		{/each}
	</ul>
</div>
