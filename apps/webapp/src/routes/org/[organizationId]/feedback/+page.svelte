<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { page } from '$app/stores';
	import type { Project } from '$lib/types';
	import type { ActionResult } from '@sveltejs/kit';
	import {
		Breadcrumb,
		BreadcrumbItem,
		Button,
		Fileupload,
		Heading,
		Helper,
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

	let importedFileRefs: any;

	let isImportModalOpen = false;

	$: feedbacks = $page.data.feedbacks;
	$: projects = $page.data.projects;

	function getProjectName(id: string) {
		return projects.find((p: Project) => p.id === id)?.name || '-';
	}

	function getTimeString(ts: number) {
		const date = new Date(ts);
		return `${date.toLocaleDateString()} ${date.toLocaleTimeString()}`;
	}

	async function uploadFeedbackFile(signedUrl: string) {
		const file = importedFileRefs[0];
		const response = await fetch(signedUrl, {
			method: 'PUT',
			mode: 'cors',
			headers: { 'content-type': file.type },
			body: file
		});
		if (!response.ok) {
			// TODO: show a friendly error instead
			throw new Error('Failed to upload file');
		}
	}

	async function processFeedbackFile(fileName: string) {
		const body = new FormData();
		body.set('name', fileName);
		const response = await fetch('?/processImportFile', { method: 'post', body });
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
	}

	async function importFeedbackFile() {
		const body = new FormData();
		const response = await fetch('?/createImportUrl', { method: 'post', body });
		const result: ActionResult = deserialize(await response.text());
		if (result.type == 'success' && result.data) {
			const { signedUrl, fileName } = result.data;
			await uploadFeedbackFile(signedUrl);
			await processFeedbackFile(fileName);
		}
		applyAction(result);
	}
</script>

<Breadcrumb class="mb-6">
	<BreadcrumbItem href={`/org/${$page.params.organizationId}`} home>Home</BreadcrumbItem>
	<BreadcrumbItem>Feedback</BreadcrumbItem>
</Breadcrumb>

<Heading customSize="text-xl font-semibold">Feedback</Heading>

{#if feedbacks.length > 0}
	<Table hoverable>
		<TableHead>
			<TableHeadCell>Created At</TableHeadCell>
			<TableHeadCell>Type</TableHeadCell>
			<TableHeadCell>Project</TableHeadCell>
			<TableHeadCell>Tags</TableHeadCell>
			<TableHeadCell>Feedback Data</TableHeadCell>
		</TableHead>
		<TableBody>
			{#each feedbacks as feedback}
				<TableBodyRow>
					<TableBodyCell>{getTimeString(feedback.createdAt)}</TableBodyCell>
					<TableBodyCell>{feedback.type}</TableBodyCell>
					<TableBodyCell>{getProjectName(feedback.projectId)}</TableBodyCell>
					<TableBodyCell>
						<code>{JSON.stringify(feedback.tags)}</code>
					</TableBodyCell>
					<TableBodyCell>
						<code>{JSON.stringify(feedback.data)}</code>
					</TableBodyCell>
				</TableBodyRow>
			{/each}
		</TableBody>
	</Table>
{/if}

<section class="mt-6">
	<Button size="xs" color="light" on:click={() => (isImportModalOpen = true)}>
		+ Import feedback
	</Button>
</section>

<Modal title="Import Feedback" bind:open={isImportModalOpen} autoclose>
	<P>TODO: add helpful text and link to documentation</P>
	<Fileupload size="sm" id="import-file" bind:files={importedFileRefs} />
	<svelte:fragment slot="footer">
		<Button size="sm" color="primary" on:click={importFeedbackFile}>Import</Button>
	</svelte:fragment>
</Modal>
