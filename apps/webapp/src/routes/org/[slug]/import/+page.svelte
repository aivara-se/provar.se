<script lang="ts">
	import { CheckIcon } from 'lucide-svelte';
	import { parser, validator, uploader } from '$lib/client/import';
	import { DashLayout } from '$lib/client/layout';
	import { toast } from '$lib/client/toast';
	import { onMount } from 'svelte';
	import route from './route.meta';

	export let data;

	interface ImportOperation {
		id: string;
		status: 'pending' | 'prepared' | 'completed';
		files: FileList | null;
		items: Record<string, string>[] | null;
		total: number;
		valid: number;
		ready: number;
	}

	let importOperation: ImportOperation = {
		id: '',
		status: 'pending',
		files: null,
		items: null,
		total: 0,
		valid: 0,
		ready: 0
	};

	onMount(() => {
		const date = new Date().toISOString().split('T')[0];
		const rand = Math.random().toString(16).substring(7);
		importOperation = {
			id: `import-${date}-${rand}`,
			status: 'pending',
			files: null,
			items: null,
			total: 0,
			valid: 0,
			ready: 0
		};
	});

	async function validateFeedbackFile() {
		if (!importOperation.files || !importOperation.files.length) {
			return;
		}
		const file = importOperation.files[0];
		const text = await file.text();
		importOperation.items = parser.parse(text);
		const result = validator.validate(importOperation.items);
		importOperation.total = result.total;
		importOperation.valid = result.valid;
	}

	async function importFeedbackFile() {
		if (!importOperation.items || !importOperation.items.length) {
			return;
		}
		for (const item of importOperation.items) {
			item['meta.import-operation'] = importOperation.id;
			try {
				await uploader.upload(data.organization.id, item);
				importOperation.ready++;
			} catch (err) {
				// console.error(err);
			}
		}
		toast('success', 'File successfully imported');
	}
</script>

<DashLayout {route}>
	<section class="prose">
		<h2>Importing feedback</h2>
		<p>
			Use the file uploader below to import user feedback from a CSV file. Ensure your file follows
			the specified format for accurate processing.
		</p>
		<h3>Feedback Types</h3>
		<ul>
			<li><code>text</code> - Basic text feedback</li>
			<li><code>cnps</code> - Net promoter score</li>
			<li><code>csat</code> - Customer satisfaction</li>
		</ul>
		<h3>Required Columns</h3>
		<ul>
			<li><code>type</code> - Feedback type (text, cnps, csat)</li>
			<li><code>time</code> - Timestamp of feedback</li>
		</ul>
		<h4>User Details</h4>
		<p>
			Columns starting with <code>user.</code> describe the user who provided the feedback. User details
			can be ("id", "name", "email", "avatar"), plus additional columns.
		</p>
		<h4>Feedback Tags</h4>
		<p>
			Columns starting with <code>tags.</code> describe data that can be used to group feedback. For
			example: these can be tags like "product", "feature", "category".
		</p>
		<h4>Extra Metadata</h4>
		<p>
			Columns starting with <code>meta.</code> describe metadata colected with the feedback. For example:
			these can be metadata like "location", "device", "browser".
		</p>
		<h3>Text Feedback</h3>
		<strong>Columns</strong>
		<ul>
			<li><code>data.response-text</code> - Plain text feedback provided by the user</li>
		</ul>
		<strong>Example</strong>
		<div class="w-80 md:w-auto overflow-x-auto">
			<table class="table table-xs">
				<thead>
					<tr>
						<th>type</th>
						<th>time</th>
						<th>data.response-text</th>
						<th>user.id</th>
						<th>tags.feature</th>
						<th>meta.request-ip</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>text</td>
						<td>2021-07-01T00:00:00Z</td>
						<td>Amazing UI!</td>
						<td>0683cfde</td>
						<td>Dashboard</td>
						<td>1.1.1.1</td>
					</tr>
					<tr>
						<td>text</td>
						<td>2021-07-01T00:00:00Z</td>
						<td>App is too slow!</td>
						<td>679e5354</td>
						<td>Onboarding</td>
						<td>8.8.8.8</td>
					</tr>
				</tbody>
			</table>
		</div>
		<h3>CNPS Feedback</h3>
		<strong>Columns</strong>
		<ul>
			<li><code>data.question-type</code> - The feedback question (rating-11p)</li>
			<li><code>data.response-data</code> - The value selected by the user</li>
			<li><code>data.response-text</code> - Plain text feedback provided by the user</li>
		</ul>
		<strong>Example</strong>
		<div class="w-80 md:w-auto overflow-x-auto">
			<table class="table table-xs">
				<thead>
					<tr>
						<th>type</th>
						<th>time</th>
						<th>data.question-type</th>
						<th>data.response-data</th>
						<th>data.response-text</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>cnps</td>
						<td>2021-07-01T00:00:00Z</td>
						<td>rating-11p</td>
						<td>8</td>
						<td></td>
					</tr>
					<tr>
						<td>cnps</td>
						<td>2021-07-01T00:00:00Z</td>
						<td>rating-11p</td>
						<td>2</td>
						<td>App is too slow!</td>
					</tr>
				</tbody>
			</table>
		</div>
		<h3>CSAT Feedback</h3>
		<strong>Columns</strong>
		<ul>
			<li><code>data.question-type</code> - The feedback question (rating-11p)</li>
			<li><code>data.response-data</code> - The value selected by the user</li>
			<li><code>data.response-text</code> - Plain text feedback provided by the user</li>
		</ul>
		<strong>Example</strong>
		<div class="w-80 md:w-auto overflow-x-auto">
			<table class="table table-xs">
				<thead>
					<tr>
						<th>type</th>
						<th>time</th>
						<th>data.question-type</th>
						<th>data.response-data</th>
						<th>data.response-text</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>csat</td>
						<td>2021-07-01T00:00:00Z</td>
						<td>rating-11p</td>
						<td>8</td>
						<td></td>
					</tr>
					<tr>
						<td>csat</td>
						<td>2021-07-01T00:00:00Z</td>
						<td>rating-11p</td>
						<td>2</td>
						<td>App is too slow!</td>
					</tr>
				</tbody>
			</table>
		</div>
		<hr />
	</section>

	<section class="mt-12">
		<div class="prose mb-8">
			<p>
				All feedbacks imported with this operation will have the metadata field
				<code>import-operation</code> set to <code>{importOperation.id}</code>.
			</p>
		</div>
		<input
			type="file"
			class="file-input file-input-sm w-full shadow"
			accept="text/csv"
			bind:files={importOperation.files}
			on:change={validateFeedbackFile}
		/>

		{#if importOperation.total}
			<table class="mt-4">
				<tr>
					<td>Validate</td>
					<td class="pl-4">
						<progress
							class="progress progress-success w-64 md:w-96"
							value={importOperation.valid}
							max={importOperation.total}
						></progress>
					</td>
					<td>
						{#if importOperation.total === importOperation.valid}
							<CheckIcon class="inline-block ml-2 w-4 h-4 text-success" />
						{/if}
					</td>
				</tr>
				<tr>
					<td>Upload</td>
					<td class="pl-4">
						<progress
							class="progress progress-success w-64 md:w-96"
							value={importOperation.ready}
							max={importOperation.total}
						></progress>
					</td>
					<td>
						{#if importOperation.total === importOperation.ready}
							<CheckIcon class="inline-block ml-2 w-4 h-4 text-success" />
						{/if}
					</td>
				</tr>
			</table>
		{/if}

		<div class="flex mt-4">
			<button class="btn btn-sm" on:click={importFeedbackFile}>Import file</button>
		</div>
	</section>
</DashLayout>
