<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { page } from '$app/stores';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Heading, Input, Label, Modal, P, Timeline, TimelineItem } from 'flowbite-svelte';
	import { CheckCircleOutline } from 'flowbite-svelte-icons';

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

<Heading customSize="mb-6 text-xl font-semibold">Getting Started</Heading>

<Timeline>
	<TimelineItem title="Create an API key" date="Step 1">
		{#if credentials.length > 0}
			<P class="mb-4">
				API keys are used to authenticate your application with Provar. Use this key to configure
				your cient application. You can revoke this key at any time from organization settings page.
			</P>
			<code class="block break-words text-sm mt-2 bg-gray-100 p-2 rounded">
				{$credentials[0].key}
			</code>
		{:else}
			<P class="mb-4">
				API keys are used to authenticate your application with Provar. Create an API key to get
				started.
			</P>
			<Button size="xs" color="light" on:click={() => (isCreateModalOpen = true)}>
				Create an API Key &nbsp;
				<CheckCircleOutline class="w-3 h-3 mr-1" />
			</Button>
		{/if}
	</TimelineItem>
	<TimelineItem title="Install client library" date="Step 2">
		<P class="mb-4">
			Use the npm module to integrate Provar into your application. You can install it using the
			following command:
		</P>
		<code class="block mb-4 break-words mt-2 bg-gray-100 p-2 rounded">
			npm install @provar/provar-js
		</code>
	</TimelineItem>
	<TimelineItem title="Collect user feedback" date="Step 3">
		<P class="mb-4">
			Configure the client package with the API key you created in step 1. You can now collect
			feedback from your users.
		</P>
		<pre class="block mb-4 break-words mt-2 bg-gray-100 p-2 rounded">
import &lbrace; ProvarClient &rbrace; from '@provar/provar-js';

const provarClient = new ProvarClient(&lbrace; apiKey: 'YOUR_API_KEY' &rbrace;);
await provarClient.sendFeedback(&lbrace; ... &rbrace;);</pre>
	</TimelineItem>
</Timeline>

<Modal title="Create API Key" bind:open={isCreateModalOpen} autoclose>
	<Label for="name" class="block mb-2">Name:</Label>
	<Input id="name" required bind:value={credentialName} />
	<svelte:fragment slot="footer">
		<Button size="xs" color="dark" on:click={createCredential}>
			Create API Key &nbsp;
			<CheckCircleOutline class="w-3 h-3 mr-1" />
		</Button>
	</svelte:fragment>
</Modal>
