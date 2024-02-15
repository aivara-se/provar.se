<script lang="ts">
	import { page } from '$app/stores';
	import { getMenuItems, type Route } from '$lib/client/routes';

	export let route: Route;

	$: items = getMenuItems()
		.filter((route) => route.sidebar?.mobile)
		.map((route) => {
			return {
				id: route.id,
				href: route.getPath($page.params),
				name: route.getName($page.params),
				icon: route.sidebar!.icon,
				active: route.isActive ? route.isActive($page.url.pathname, $page.params) : false
			};
		});
</script>

<nav class="fixed left-0 right-0 bottom-0 btm-nav btm-nav-sm shadow-lg border-t dark:border-t-0">
	{#each items as item}
		<a href={item.href}>
			<svelte:component this={item.icon} class="w-4 h-4" />
			<span class="btm-nav-label">{item.name}</span>
		</a>
	{/each}
</nav>
