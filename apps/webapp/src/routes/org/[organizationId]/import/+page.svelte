<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { DashLayout } from '$lib/client/layout';
	import type { ActionResult } from '@sveltejs/kit';
	import route from './route.meta';
	import { toast } from '../../../../lib/client/toast';

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
		await applyAction(result);
		toast('success', 'File successfully imported');
	}
</script>

<DashLayout {route}>
	<div class="">
		<input type="file" class="file-input w-full shadow" bind:files={importedFileRefs} />
		<div class="flex mt-4">
			<button class="btn" on:click={importFeedbackFile}>Import file</button>
		</div>
	</div>
</DashLayout>
