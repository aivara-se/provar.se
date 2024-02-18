<script lang="ts">
	import { page } from '$app/stores';
	import { getMenuItems } from '$lib/client/routes';

	$: items = getMenuItems()
		.filter((route) => route.sidebar?.mobile)
		.map((route) => {
			return {
				id: route.id,
				icon: route.icon,
				href: route.getPath($page.params),
				name: route.getName($page.params),
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
