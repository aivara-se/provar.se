<script lang="ts">
	import type { Feedback } from '$lib/types';
	import { Accordion, AccordionItem, Badge, Blockquote, P } from 'flowbite-svelte';
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import { FileCodeOutline, QuoteSolid, TagOutline } from 'flowbite-svelte-icons';
	import FeedbackUserAvatar from './_components/FeedbackUserAvatar.svelte';

	export let feedback: Feedback | null = null;

	$: feedbackJSON = JSON.stringify(feedback, null, 2);
</script>

<section class="content">
	<div class="content-scroll">
		<div class="h-0">
			{#if !feedback}
				<section>
					<P class="text-gray-500 text-sm">Select a feedback from the list to see details</P>
				</section>
			{/if}

			{#if feedback}
				<section>
					<Blockquote bg border baseClass="p-2 px-4">
						<div class="flex justify-end"><QuoteSolid class="w-5 h-5 text-gray-400" /></div>
						“{feedback.data.text || 'No feedback text'}”
						<FeedbackUserAvatar user={feedback.user} withName={true} />
					</Blockquote>
				</section>

				<section class="mt-3 flex gap-2">
					{#each Object.keys(feedback.tags || {}).sort() as key}
						<Badge border color="dark" class="border-gray-300 text-gray-600">
							<TagOutline class="w-2.5 h-2.5 me-1.5" />
							{key}: {feedback.tags?.[key]}
						</Badge>
					{/each}
				</section>

				<section>
					<Accordion flush>
						<AccordionItem open>
							<span slot="header">
								<div class="flex items-center">Feedback Metadata</div>
							</span>
							<Table>
								<TableHead>
									<TableHeadCell>Key</TableHeadCell>
									<TableHeadCell>Value</TableHeadCell>
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
										</TableBodyRow>
									{/each}
								</TableBody>
							</Table>
						</AccordionItem>
					</Accordion>
				</section>
			{/if}
		</div>
	</div>
</section>

<style>
	.content {
		display: flex;
		flex: 1;
		min-height: 0px;
		padding: 20px 0 0 20px;
		box-sizing: border-box;
	}

	.content-scroll {
		display: flex;
		flex: 1;
		flex-direction: column;
		overflow-x: hidden;
		overflow-y: auto;
	}

	.content-scroll::-webkit-scrollbar {
		display: none;
	}
</style>
