<script lang="ts">
	import { page } from '$app/stores';
	import { BookIcon } from 'lucide-svelte';
	import { getMenuItems, type Route } from '$lib/client/routes';

	export let route: Route;

	$: items = getMenuItems().map((route) => {
		return {
			id: route.id,
			icon: route.icon || BookIcon,
			href: route.getPath($page.params),
			name: route.getName($page.params),
			active: route.isActive ? route.isActive($page.url.pathname, $page.params) : false
		};
	});
</script>

<ul class="menu p-0 w-32 rounded-box">
	{#each items as item}
		<li>
			<a href={item.href} class="active:bg-transparent">
				{#if item.icon}
					<svelte:component
						this={item.icon}
						class="w-4 h-4 {item.active ? 'opacity-100' : 'opacity-50'}"
					/>
				{/if}
				<span class="antialiased text-sm {item.active ? 'opacity-100' : 'opacity-50'}">
					{item.name}
				</span>
			</a>
		</li>
	{/each}
</ul>
