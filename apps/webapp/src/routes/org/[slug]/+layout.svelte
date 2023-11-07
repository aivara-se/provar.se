<script lang="ts">
	import { page } from '$app/stores';
	import { LSidebarLayout } from '$lib/ui';
	import {
		AdjustmentsHorizontalOutline,
		ChartMixedOutline,
		MessageDotsOutline,
		MessagesOutline,
		UserSettingsOutline
	} from 'flowbite-svelte-icons';
	import type { LayoutData } from './$types';
	import OrganizationSelector from './_components/OrganizationSelector.svelte';
	import SidebarNavigationMenu from './_components/SidebarNavigationMenu.svelte';
	import UserActionButtonGroup from './_components/UserActionButtonGroup.svelte';

	export let data: LayoutData;

	/**
	 * Navigation items
	 */
	$: navigationItems = [
		{ name: 'Dashboard', href: `/org/${$page.params.slug}`, icon: ChartMixedOutline },
		{ name: 'Feedbacks', href: `/org/${$page.params.slug}/feedbacks`, icon: MessagesOutline }
	].map((item) => ({
		...item,
		active: item.href === $page.url.pathname
	}));

	/**
	 * User action items
	 */
	$: userActionItems = [
		{ name: 'Profile', href: `/auth/logout`, icon: UserSettingsOutline },
		{
			name: 'Preferences',
			href: `/org/${$page.params.slug}/preferences`,
			icon: AdjustmentsHorizontalOutline
		},
		{ name: 'Chat UI', href: `/org/${$page.params.slug}/chat`, icon: MessageDotsOutline }
	].map((item) => ({
		...item,
		active: item.href === $page.url.pathname
	}));
</script>

<LSidebarLayout>
	<svelte:fragment slot="sidebar-top">
		<SidebarNavigationMenu items={navigationItems} />
	</svelte:fragment>

	<svelte:fragment slot="sidebar-bottom">
		<OrganizationSelector value={$page.params.slug} items={data.organizations} />
		<UserActionButtonGroup items={userActionItems} />
	</svelte:fragment>

	<svelte:fragment>
		<slot />
	</svelte:fragment>
</LSidebarLayout>
