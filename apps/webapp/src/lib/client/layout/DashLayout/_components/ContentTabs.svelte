<script lang="ts">
	import { page } from '$app/stores';
	import { getRoute } from '$lib/client/routes';

	export let tabs: string[] = [];

	$: items = tabs.map((id) => {
		const route = getRoute(id);
		return {
			id: route.id,
			icon: route.icon,
			href: route.getPath($page.params),
			name: route.getName($page.params),
			active: route.isActive ? route.isActive($page.url.pathname, $page.params) : false
		};
	});
</script>

<div role="tablist" class="tabs tabs-boxed mb-4">
	{#each items as item}
		<a role="tab" href={item.href} class="tab {item.active ? 'tab-active' : ''}">
			{#if item.icon}
				<svelte:component this={item.icon} class="w-4 h-4" />
			{/if}
			{item.name}
		</a>
	{/each}
</div>
