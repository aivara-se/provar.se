<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import { toast } from '$lib/client/toast';
	import route from './route.meta';

	let importedFileRefs: FileList;

	async function importFeedbackFile() {
		// TODO: call api and create signed import url
		// TODO: upload the file to new the signed url
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
