<script lang="ts">
	import { page } from '$app/stores';
	import { HorizontalSeparator, LSidebarLayout } from '$lib/ui';
	import { Button, ButtonGroup } from 'flowbite-svelte';
	import {
		AdjustmentsHorizontalOutline,
		UserSettingsOutline,
		MessageDotsOutline,
		ChartMixedOutline,
		MessagesOutline
	} from 'flowbite-svelte-icons';
	import type { LayoutData } from './$types';
	import OrganizationSelector from './_components/OrganizationSelector.svelte';
	import SidebarNavigationMenu from './_components/SidebarNavigationMenu.svelte';

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
</script>

<LSidebarLayout>
	<svelte:fragment slot="sidebar-top">
		<SidebarNavigationMenu items={navigationItems} />
	</svelte:fragment>

	<svelte:fragment slot="sidebar-bottom">
		<OrganizationSelector value={$page.params.slug} items={data.organizations} />

		<ButtonGroup class="mt-2">
			<Button href={`/auth/logout`}>
				<UserSettingsOutline class="w-4 h-4 mr-1" />
			</Button>
			<Button href={`/org/${$page.params.slug}/preferences`}>
				<AdjustmentsHorizontalOutline class="w-4 h-4 mr-1" />
			</Button>
			<Button>
				<MessageDotsOutline class="w-4 h-4 mr-1" />
			</Button>
		</ButtonGroup>
	</svelte:fragment>

	<svelte:fragment>
		<slot />
	</svelte:fragment>
</LSidebarLayout>
