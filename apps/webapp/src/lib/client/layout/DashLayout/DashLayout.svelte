<script lang="ts">
	import type { Route } from '$lib/client/routes';
	import AvatarMenu from './_components/AvatarMenu.svelte';
	import Breadcrumbs from './_components/Breadcrumbs.svelte';
	import ContentTabs from './_components/ContentTabs.svelte';
	import MobileMenu from './_components/MobileMenu.svelte';
	import SidebarMenu from './_components/SidebarMenu.svelte';

	export let route: Route;
</script>

<div class="flex flex-1 min-h-screen w-full h-full bg-gray-100 dark:bg-gray-900 pb-16 md:pb-4">
	<aside class="hidden md:flex flex-col w-32 ml-4 relative">
		<div class="fixed flex flex-col items-center top-12 bottom-0 left-4 w-32">
			<div class="flex flex-1">
				<SidebarMenu {route} />
			</div>
		</div>
	</aside>

	<section class="flex-1 mx-4">
		<header class="flex flex-row items-center justify-between h-12 px-2">
			<Breadcrumbs {route} />
			<AvatarMenu />
		</header>
		<main class="w-full min-h-screen p-4 md:p-8 rounded-xl bg-white dark:bg-black/20 shadow">
			{#if route.tabs && route.tabs.length > 0}
				<ContentTabs tabs={route.tabs} />
			{/if}

			<slot />
		</main>
	</section>

	<nav class="flex md:hidden">
		<MobileMenu {route} />
	</nav>
</div>
