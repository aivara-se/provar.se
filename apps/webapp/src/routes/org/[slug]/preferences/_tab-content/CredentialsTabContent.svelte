<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { page } from '$app/stores';
	import type { ActionResult } from '@sveltejs/kit';
	import {
		A,
		Button,
		Heading,
		Input,
		Label,
		Modal,
		P,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import { CheckCircleOutline } from 'flowbite-svelte-icons';

	let name = '';

	let selectedCredential: any;

	let isCreateModalOpen = false;
	let isDetailsModalOpen = false;

	$: credentials = $page.data.credentials;

	function selectCredential(cred: any) {
		selectedCredential = cred;
		isDetailsModalOpen = true;
	}

	async function createCredential() {
		const data = new FormData();
		data.set('name', name);
		const response = await fetch('?/createCredential', {
			method: 'post',
			body: data
		});
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		name = '';
		applyAction(result);
	}

	async function revokeCredential(id: string) {
		const data = new FormData();
		data.set('id', id);
		const response = await fetch('?/revokeCredential', {
			method: 'post',
			body: data
		});
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
	}
</script>

<section class="create-credentials">
	<Heading customSize="text-xl font-semibold">API Keys</Heading>

	<P class="mt-6">
		API keys are used to authenticate your application with the API. These keys provides your
		application access to our platform's features.
	</P>

	<section class="create-credentials mt-6">
		<Button size="xs" color="light" on:click={() => (isCreateModalOpen = true)}>
			Create an API Key &nbsp;
			<CheckCircleOutline class="w-3 h-3 mr-1" />
		</Button>
	</section>

	{#if credentials.length > 0}
		<section class="current-credentials mt-6">
			<Table hoverable>
				<TableHead>
					<TableHeadCell>Credential name</TableHeadCell>
					<TableHeadCell>Created At</TableHeadCell>
					<TableHeadCell>Last Used</TableHeadCell>
					<TableHeadCell />
				</TableHead>
				<TableBody>
					{#each credentials as item}
						<TableBodyRow>
							<TableBodyCell>{item.name}</TableBodyCell>
							<TableBodyCell>-</TableBodyCell>
							<TableBodyCell>-</TableBodyCell>
							<TableBodyCell>
								<Button pill size="xs" color="light" on:click={() => selectCredential(item)}>
									Details
								</Button>
							</TableBodyCell>
						</TableBodyRow>
					{/each}
				</TableBody>
			</Table>
		</section>
	{/if}
</section>

<Modal title="Create API Key" bind:open={isCreateModalOpen} autoclose>
	<Label for="name" class="block mb-2">Name:</Label>
	<Input id="name" required bind:value={name} />
	<svelte:fragment slot="footer">
		<Button size="xs" color="dark" on:click={createCredential}>
			Create API Key &nbsp;
			<CheckCircleOutline class="w-3 h-3 mr-1" />
		</Button>
	</svelte:fragment>
</Modal>

<Modal title={selectedCredential?.name} bind:open={isDetailsModalOpen} autoclose>
	<P class="mb-2">
		<span class="font-semibold">Client ID:</span>
		<code class="block break-words text-xs mt-2 bg-gray-100 p-2 rounded">
			{selectedCredential?.clientId}
		</code>
	</P>
	<P class="mb-2">
		<span class="font-semibold">Client Secret:</span>
		<code class="block break-words text-xs mt-2 bg-gray-100 p-2 rounded">
			{selectedCredential?.clientSecret}
		</code>
	</P>
	<svelte:fragment slot="footer">
		<Button size="xs" color="red" on:click={() => revokeCredential(selectedCredential.id)}>
			Revoke
		</Button>
	</svelte:fragment>
</Modal>
