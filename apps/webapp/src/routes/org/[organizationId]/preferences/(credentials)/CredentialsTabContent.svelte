<script lang="ts">
	import { page } from '$app/stores';
	import type { Credential } from '$lib/types';
	import {
		Button,
		Heading,
		P,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import CreateCredentialModal from './_components/CreateCredentialModal.svelte';
	import CredentialDetailsModal from './_components/CredentialDetailsModal.svelte';

	let selectedCredential: Credential;

	let isCreateModalOpen = false;
	let isDetailsModalOpen = false;

	$: credentials = $page.data.credentials || [];

	function selectCredential(cred: Credential) {
		selectedCredential = cred;
		isDetailsModalOpen = true;
	}
</script>

<section>
	<Heading customSize="text-xl font-semibold">API Keys</Heading>

	<P class="mt-6">
		API keys are used to authenticate your application with the API. These keys allow your
		application to upload feedback data.
	</P>

	{#if credentials.length > 0}
		<section class="current-credentials mt-6">
			<Table hoverable>
				<TableHead>
					<TableHeadCell>API key name</TableHeadCell>
					<TableHeadCell>Created At</TableHeadCell>
					<TableHeadCell>Last Used</TableHeadCell>
					<TableHeadCell />
				</TableHead>
				<TableBody>
					{#each credentials as item}
						<TableBodyRow>
							<TableBodyCell>{item.name}</TableBodyCell>
							<TableBodyCell>{item.createdAt.toLocaleDateString()}</TableBodyCell>
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

<section class="mt-6">
	<Button size="xs" color="light" on:click={() => (isCreateModalOpen = true)}>
		+ Create API Key
	</Button>
</section>

<CreateCredentialModal bind:isOpen={isCreateModalOpen} />
{#if selectedCredential}
	<CredentialDetailsModal bind:isOpen={isDetailsModalOpen} bind:credential={selectedCredential} />
{/if}
