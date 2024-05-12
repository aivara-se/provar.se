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

<div role="tablist" class="tabs tabs-bordered tabs-sm mb-4 bg-transparent">
	{#each items as item}
		<a role="tab" href={item.href} class="tab text-sm pb-2 {item.active ? 'tab-active' : ''}">
			{#if item.icon}
				<svelte:component this={item.icon} class="w-4 h-4" />
			{/if}
		</a>
	{/each}
</div>
