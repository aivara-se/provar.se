<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import type { ActionResult } from '@sveltejs/kit';
	import { Button, Dropdown, DropdownItem, Fileupload, Modal, P } from 'flowbite-svelte';
	import {
		ArrowDownToBracketOutline,
		ArrowUpFromBracketOutline,
		DotsHorizontalOutline
	} from 'flowbite-svelte-icons';

	let isExportModalOpen = false;
	let isImportModalOpen = false;
	let importedFileRefs: FileList;

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

<div>
	<Button size="xs" color="light">
		<span>&nbsp;</span>
		<DotsHorizontalOutline class="w-3 h-3 text-gray-500 pointer-events-none" />
	</Button>
	<Dropdown>
		<DropdownItem class="flex items-center" on:click={() => (isExportModalOpen = true)}>
			<ArrowDownToBracketOutline class="w-3 h-3 mr-2" />
			Export
		</DropdownItem>
		<DropdownItem class="flex items-center" on:click={() => (isImportModalOpen = true)}>
			<ArrowUpFromBracketOutline class="w-3 h-3 mr-2" />
			Import
		</DropdownItem>
	</Dropdown>
</div>

<Modal title="Export Feedback" bind:open={isExportModalOpen} autoclose>
	<P>TODO: add helpful text and link to documentation</P>
	<svelte:fragment slot="footer">
		<Button size="sm" color="primary" on:click={() => {}}>Export</Button>
	</svelte:fragment>
</Modal>

<Modal title="Import Feedback" bind:open={isImportModalOpen} autoclose>
	<P>TODO: add helpful text and link to documentation</P>
	<Fileupload size="sm" id="import-file" bind:files={importedFileRefs} />
	<svelte:fragment slot="footer">
		<Button size="sm" color="primary" on:click={importFeedbackFile}>Import</Button>
	</svelte:fragment>
</Modal>
