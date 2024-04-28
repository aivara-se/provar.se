<script lang="ts">
	import { page } from '$app/stores';
	import { getMenuItems, getRoute } from '$lib/client/routes';
	import { BookIcon } from 'lucide-svelte';

	$: items = getMenuItems().map((route) => {
		return {
			id: route.id,
			icon: route.icon || BookIcon,
			href: route.getPath($page.params),
			name: route.getName($page.params),
			active: route.isActive ? route.isActive($page.url.pathname, $page.params) : false,
			kids: route.tabs ? route.tabs.map(formatChildRoute) : []
		};
	});

	function formatChildRoute(id: string) {
		const route = getRoute(id);
		return {
			id: route.id,
			icon: route.icon,
			href: route.getPath($page.params),
			name: route.getName($page.params),
			active: route.isActive ? route.isActive($page.url.pathname, $page.params) : false
		};
	}
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
			{#if item.active && item.kids.length > 0}
				<ul>
					{#each item.kids as kid}
						<li>
							<a href={kid.href}>
								<span class="antialiased text-xs {kid.active ? 'opacity-100' : 'opacity-50'}">
									{kid.name}
								</span>
							</a>
						</li>
					{/each}
				</ul>
			{/if}
		</li>
	{/each}
</ul>
