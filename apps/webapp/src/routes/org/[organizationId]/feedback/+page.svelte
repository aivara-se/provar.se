<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { goto, invalidateAll } from '$app/navigation';
	import { page } from '$app/stores';
	import type { ActionResult } from '@sveltejs/kit';
	import {
		Breadcrumb,
		BreadcrumbItem,
		Button,
		ButtonGroup,
		Dropdown,
		DropdownItem,
		Fileupload,
		Heading,
		Input,
		Modal,
		P,
		Pagination,
		type LinkType,
		Card,
		Badge
	} from 'flowbite-svelte';
	import {
		ArrowDownToBracketOutline,
		ArrowUpFromBracketOutline,
		ChevronLeftOutline,
		ChevronRightOutline,
		DotsHorizontalOutline,
		SearchOutline
	} from 'flowbite-svelte-icons';
	import { stringify } from 'postcss';

	let isExportModalOpen = false;
	let isImportModalOpen = false;
	let importedFileRefs: FileList;

	let pages: LinkType[] = [];
	$: currentPage = Number.parseInt($page.url.searchParams.get('page') || '1');

	$: {
		pages = [];
		const min = Math.max(1, currentPage - 2);
		const max = Math.min($page.data.feedbacks.pages, currentPage + 2);
		for (let i = min; i <= max; i++) {
			const name =
				i === min && min > 1
					? '...'
					: i === max && max < $page.data.feedbacks.pages
					  ? '...'
					  : i.toString();
			pages.push({
				name,
				href: `?page=${i}`,
				active: i === currentPage
			});
		}
	}

	function previousPage() {
		if (currentPage === 1) {
			return;
		}
		goto(`?page=${currentPage - 1}`);
	}

	function nextPage() {
		if (currentPage === $page.data.feedbacks.pages) {
			return;
		}
		goto(`?page=${currentPage + 1}`);
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

<section class="flex">
	<Heading class="flex-1" customSize="text-xl font-semibold">Feedback</Heading>
	<Button size="xs" color="light">
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
</section>

<section class="flex flex-1">
	<aside class="sidebar">
		<section class="sidebar-top">
			<div>
				<ButtonGroup class="w-full">
					<Input id="search" size="sm" />
					<Button color="dark">
						<SearchOutline class="w-3 h-3" />
					</Button>
				</ButtonGroup>
			</div>
		</section>
		<section class="sidebar-main">
			<div class="sidebar-main-scroll">
				<div class="h-0">
					{#each $page.data.feedbacks.items as feedback}
						<Card shadow={false} padding="none" class="mb-6 py-2 px-3">
							<div class="flex justify-end">
								<Badge color="none">{feedback.type}</Badge>
							</div>
							{#if feedback.data.text}
								<P size="sm">“{feedback.data.text}”</P>
							{:else}
								<P size="sm" italic class="text-gray-500">No feedback text</P>
							{/if}
							<P size="xs" class="mt-2 text-gray-500">
								{feedback.createdAt.toLocaleDateString()} at {feedback.createdAt.toLocaleTimeString()}
							</P>
						</Card>
					{/each}
				</div>
			</div>
		</section>
		<section class="sidebar-bottom">
			<Pagination {pages} on:previous={previousPage} on:next={nextPage} icon>
				<svelte:fragment slot="prev">
					<span class="sr-only">Previous</span>
					<ChevronLeftOutline class="w-2 h-2 outline-none" />
				</svelte:fragment>
				<svelte:fragment slot="next">
					<span class="sr-only">Next</span>
					<ChevronRightOutline class="w-2 h-2 outline-none" />
				</svelte:fragment>
			</Pagination>
		</section>
	</aside>
	<main class="content">DETAILS</main>
</section>

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

<style>
	.sidebar {
		display: flex;
		padding: 20px 20px 0 0;
		width: 360px;
		flex-direction: column;
		box-sizing: border-box;
		border-right: solid 1px #e0e0e0;
	}

	.sidebar-top {
		display: flex;
		flex-direction: column;
	}

	.sidebar-main {
		display: flex;
		flex: 1;
		min-height: 0px;
		padding: 20px 0 0 0;
		box-sizing: border-box;
	}

	.sidebar-main-scroll {
		display: flex;
		flex: 1;
		flex-direction: column;
		overflow-x: hidden;
		overflow-y: auto;
	}

	.sidebar-main-scroll::-webkit-scrollbar {
		display: none;
	}

	.sidebar-bottom {
		display: flex;
		justify-content: center;
		padding: 20px 0 0 0;
		box-sizing: border-box;
	}

	.content {
		flex: 1;
		display: flex;
		flex-direction: column;
		padding: 20px 0 0 20px;
		box-sizing: border-box;
	}
</style>
