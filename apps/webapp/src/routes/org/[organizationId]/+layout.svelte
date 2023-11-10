<script lang="ts">
	import { page } from '$app/stores';
	import { LSidebarLayout } from '$lib/ui';
	import { sendTextFeedback } from '$lib/provar';
	import { Button, Modal, Textarea } from 'flowbite-svelte';
	import {
		AdjustmentsHorizontalOutline,
		ChartMixedOutline,
		CheckCircleOutline,
		HomeOutline,
		MessageDotsOutline,
		MessagesOutline,
		UserSettingsOutline
	} from 'flowbite-svelte-icons';
	import type { LayoutData } from './$types';
	import OrganizationSelector from './_components/OrganizationSelector.svelte';
	import SidebarNavigationMenu from './_components/SidebarNavigationMenu.svelte';
	import UserActionButtonGroup from './_components/UserActionButtonGroup.svelte';

	export let data: LayoutData;

	let feedbackMessage = '';
	let feedbackModalVisible = false;

	async function sendUserFeedback() {
		await sendTextFeedback(feedbackMessage);
		feedbackMessage = '';
		feedbackModalVisible = false;
	}

	/**
	 * Navigation items
	 */
	$: navigationItems = [
		{
			name: 'Home',
			href: `/org/${$page.params.organizationId}`,
			icon: HomeOutline
		},
		{
			name: 'Overview',
			href: `/org/${$page.params.organizationId}/overview`,
			icon: ChartMixedOutline
		},
		{
			name: 'Feedback',
			href: `/org/${$page.params.organizationId}/feedback`,
			icon: MessagesOutline
		}
	].map((item) => ({
		...item,
		active: item.href === $page.url.pathname
	}));

	/**
	 * User action items
	 */
	$: userActionItems = [
		{
			name: 'Profile',
			href: `/auth/logout`,
			icon: UserSettingsOutline
		},
		{
			name: 'Preferences',
			href: `/org/${$page.params.organizationId}/preferences`,
			icon: AdjustmentsHorizontalOutline
		},
		{
			name: 'Chat UI',
			onClick: () => (feedbackModalVisible = true),
			active: feedbackModalVisible,
			icon: MessageDotsOutline
		}
	].map((item) => ({
		...item,
		active: item.active || item.href === $page.url.pathname
	}));
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

<Modal title="Send Feedback" bind:open={feedbackModalVisible} autoclose>
	<Textarea
		id="message"
		placeholder="Your message"
		rows="4"
		name="message"
		required
		bind:value={feedbackMessage}
	/>
	<svelte:fragment slot="footer">
		<Button size="xs" color="primary" on:click={sendUserFeedback}>
			Send Feedback &nbsp;
			<CheckCircleOutline class="w-3 h-3 mr-1" />
		</Button>
	</svelte:fragment>
</Modal>
