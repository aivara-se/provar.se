<script lang="ts">
	import { mergeSearchValue } from '$lib/client/search';
	import type { Feedback } from '$lib/types';
	import {
		Accordion,
		AccordionItem,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import { FilterOutline } from 'flowbite-svelte-icons';

	export let feedback: Feedback;

	function filterByMetadata(key: string, val: unknown) {
		mergeSearchValue({ meta: { [key]: val } });
	}
</script>

<section class="mt-3">
	<Accordion flush>
		<AccordionItem open>
			<span slot="header">
				<div class="flex items-center">Feedback Metadata</div>
			</span>
			<Table>
				<TableHead>
					<TableHeadCell>Key</TableHeadCell>
					<TableHeadCell>Value</TableHeadCell>
					<TableHeadCell></TableHeadCell>
				</TableHead>
				<TableBody>
					{#each Object.keys(feedback.meta || {}).sort() as key}
						<TableBodyRow>
							<TableBodyCell tdClass="px-6 py-4 font-medium w-64">
								{key}
							</TableBodyCell>
							<TableBodyCell tdClass="px-6 py-4 font-normal">
								{feedback.meta?.[key]}
							</TableBodyCell>
							<TableBodyCell tdClass="px-6 py-4 font-normal">
								<FilterOutline
									class="w-3 h-3 text-gray-500 cursor-pointer"
									on:click={() => filterByMetadata(key, feedback.meta?.[key])}
								/>
							</TableBodyCell>
						</TableBodyRow>
					{/each}
				</TableBody>
			</Table>
		</AccordionItem>
	</Accordion>
</section>
