<script lang="ts">
	import { page } from '$app/stores';
	import { LSidebarLayout, HorizontalSeparator } from '$lib/ui';
	import OrganizationSelector from './_components/OrganizationSelector.svelte';
	import SidebarNavigationMenu from './_components/SidebarNavigationMenu.svelte';

	/**
	 * Navigation items
	 */
	$: navigationItems = [
		{ name: 'Dashboard', href: '/org/' + $page.params.slug },
		{ name: 'Projects', href: '/org/' + $page.params.slug + '/projects' },
		{ name: 'Feedback', href: '/org/' + $page.params.slug + '/feedback' },
		{ name: 'Features', href: '/org/' + $page.params.slug + '/features' }
	].map((item) => ({
		...item,
		active: item.href === $page.url.pathname
	}));
</script>

<LSidebarLayout>
	<svelte:fragment slot="sidebar-top">
		<OrganizationSelector value={$page.params.slug} items={$page.data.organizations} />
		<HorizontalSeparator />
		<SidebarNavigationMenu items={navigationItems} />
	</svelte:fragment>

	<svelte:fragment slot="sidebar-bottom">logo</svelte:fragment>

	<svelte:fragment>
		<slot />
	</svelte:fragment>
</LSidebarLayout>
