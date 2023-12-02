<script lang="ts">
	import { page } from '$app/stores';
	import type { Project } from '$lib/types';
	import {
		Breadcrumb,
		BreadcrumbItem,
		Heading,
		P,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';

	$: feedbacks = $page.data.feedbacks;
	$: projects = $page.data.projects;

	function getProjectName(id: string) {
		return projects.find((p: Project) => p.id === id)?.name || '-';
	}

	function getTimeString(ts: number) {
		const date = new Date(ts);
		return `${date.toLocaleDateString()} ${date.toLocaleTimeString()}`;
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
