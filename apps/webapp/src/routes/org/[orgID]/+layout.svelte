<script lang="ts">
	import { page } from '$app/stores';
	import { OrganizationSelector, MainNavigationMenu, SidebarSeparator } from '$lib/ui';

	/**
	 * Selected organization information (id, name, etc.)
	 */
	$: selectedOrganization = $page.data.organizations.find(
		(org: any) => org.id === $page.params.orgID
	);

	/**
	 * Navigation items
	 */
	$: navigationItemsList = [
		{ name: 'Dashboard', href: '/org/' + $page.params.orgID },
		{ name: 'Projects', href: '/org/' + $page.params.orgID + '/projects' },
		{ name: 'Feedback', href: '/org/' + $page.params.orgID + '/feedback' },
		{ name: 'Features', href: '/org/' + $page.params.orgID + '/features' }
	].map((item) => ({
		...item,
		active: item.href === $page.url.pathname
	}));
</script>

<div class="wrapper">
	<aside class="sidebar">
		<section class="sidebar-top">
			<OrganizationSelector
				organization={selectedOrganization}
				organizations={$page.data.organizations}
			/>
			<SidebarSeparator />
			<MainNavigationMenu items={navigationItemsList} />
		</section>
		<section class="sidebar-bottom">logo</section>
	</aside>
	<main class="content"><slot /></main>
</div>

<style>
	.wrapper {
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: center;
		height: 100vh;
	}

	.sidebar {
		display: flex;
		padding: 20px;
		width: 200px;
		height: 100%;
		background: #f0f0f0;
		flex-direction: column;
		box-sizing: border-box;
	}

	.sidebar-top {
		display: flex;
		flex: 1;
		flex-direction: column;
	}

	.sidebar-bottom {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.content {
		display: flex;
		flex: 1;
		height: 100%;
		flex-direction: column;
		box-sizing: border-box;
		padding: 20px;
	}
</style>
