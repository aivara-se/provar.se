<script lang="ts">
	import { page } from '$app/stores';
	import { LSidebarLayout } from '$lib/client/ui';
	import { signOut } from '@auth/sveltekit/client';
	import {
		AdjustmentsHorizontalOutline,
		ChartMixedOutline,
		ClockOutline,
		MessageDotsOutline,
		UserHeadsetOutline,
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
		{
			name: 'Feedback',
			href: `/org/${$page.params.organizationId}/feedback`,
			icon: ChartMixedOutline
		},
		{
			name: 'Projects',
			href: `/org/${$page.params.organizationId}/project`,
			icon: ClockOutline
		},
		{
			name: 'Settings',
			href: `/org/${$page.params.organizationId}/preferences`,
			icon: AdjustmentsHorizontalOutline
		}
	].map((item) => ({
		...item,
		active: $page.url.pathname.startsWith(item.href)
	}));

	/**
	 * User action items
	 */
	$: userActionItems = [
		{
			name: 'Profile',
			onClick: () => signOut(),
			icon: UserSettingsOutline
		},
		{
			name: 'Chat UI',
			onClick: () => alert('Chat UI'),
			icon: MessageDotsOutline
		},
		{
			name: 'Support',
			onClick: () => alert('Support'),
			icon: UserHeadsetOutline
		}
	];
</script>

<LSidebarLayout>
	<svelte:fragment slot="sidebar-top">
		<SidebarNavigationMenu items={navigationItems} />
	</svelte:fragment>

	<svelte:fragment slot="sidebar-bottom">
		<OrganizationSelector current={$page.params.organizationId} items={data.organizations} />
		<UserActionButtonGroup items={userActionItems} />
	</svelte:fragment>

	<svelte:fragment>
		<slot />
	</svelte:fragment>
</LSidebarLayout>
