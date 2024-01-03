<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { page } from '$app/stores';
	import type { ActionResult } from '@sveltejs/kit';
	import {
		Breadcrumb,
		BreadcrumbItem,
		Button,
		Heading,
		Input,
		Label,
		Modal,
		P,
		Timeline,
		TimelineItem
	} from 'flowbite-svelte';
	import { CopyableCode } from '../../../../lib/client/ui';

	$: credentials = $page.data.credentials || [];

	let credentialName: string;
	let isCreateModalOpen: boolean;

	async function createCredential() {
		const data = new FormData();
		data.set('name', credentialName);
		const action = `/org/${$page.params.organizationId}/preferences?/createCredential`;
		const response = await fetch(action, {
			method: 'post',
			body: data
		});
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		credentialName = '';
		applyAction(result);
	}
</script>

<Breadcrumb class="mb-6">
	<BreadcrumbItem href={`/org/${$page.params.organizationId}`} home>Home</BreadcrumbItem>
	<BreadcrumbItem>Getting Started</BreadcrumbItem>
</Breadcrumb>

<Heading customSize="mb-2 text-xl font-semibold">Getting started</Heading>

<Timeline>
	<TimelineItem title="Create an API key" date="Step 1">
		<section class="mt-2">
			{#if credentials.length > 0}
				<P class="mb-4">
					API keys are used to authenticate your application with Provar. Use this key to configure
					your cient application. You can revoke this key at any time from organization settings
					page.
				</P>
				<CopyableCode text={credentials[0].key} />
			{:else}
				<P class="mb-4">
					API keys are used to authenticate your application with Provar. Create an API key to get
					started.
				</P>
				<Button size="xs" color="light" on:click={() => (isCreateModalOpen = true)}>
					+ Create API Key
				</Button>
			{/if}
		</section>
	</TimelineItem>
	<TimelineItem title="Install client library" date="Step 2">
		<section class="mt-2">
			<P class="mb-4">Install the @provar/provar-js client library</P>
			<CopyableCode text="npm install @provar/provar-js" />
		</section>
	</TimelineItem>
	<TimelineItem title="Collect user feedback" date="Step 3">
		<section class="mt-2">
			<P class="mb-4">Configure the client package with the API key.</P>
			<CopyableCode
				text={`
import { ProvarClient } from '@provar/provar-js';

const provarClient = new ProvarClient({ apiKey: 'YOUR_API_KEY' });
await provarClient.sendText('My feedback');`}
			/>
			<P class="mt-4">You can now collect feedback from your users.</P>
		</section>
	</TimelineItem>
</Timeline>

<Modal title="Create API Key" bind:open={isCreateModalOpen} autoclose>
	<Label for="name" class="block mb-2">Name:</Label>
	<Input id="name" required bind:value={credentialName} />
	<svelte:fragment slot="footer">
		<Button size="sm" color="primary" on:click={createCredential}>Create</Button>
	</svelte:fragment>
</Modal>
