<script lang="ts">
	import { OverlayScrollbarsComponent } from 'overlayscrollbars-svelte';
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
	<section class="prose">
		<h2>Importing feedback</h2>
		<p>
			Use the file uploader below to import user feedback from a CSV file. Ensure your file follows
			the specified format for accurate processing.
		</p>
		<h3>Feedback Types</h3>
		<ul>
			<li><code>text</code> - Plain text feedback</li>
			<li><code>cnps</code> - Customer net promoter score feedback</li>
			<li><code>csat</code> - Customer satisfaction score feedback</li>
		</ul>
		<h3>Required Columns</h3>
		<ul>
			<li><code>type</code> - The feedback type ("text", "cnps", "csat")</li>
			<li><code>time</code> - Timestamp of feedback in ISO format</li>
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
		<h3>CNPS Feedback</h3>
		<strong>Columns</strong>
		<ul>
			<li><code>data.question-type</code> - The feedback question ("rating-11p")</li>
			<li><code>data.response-data</code> - The value selected by the user</li>
			<li><code>data.response-text</code> - Plain text feedback provided by the user</li>
		</ul>
		<strong>Example</strong>
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
		<h3>CSAT Feedback</h3>
		<strong>Columns</strong>
		<ul>
			<li><code>data.question-type</code> - The feedback question ("rating-11p")</li>
			<li><code>data.response-data</code> - The value selected by the user</li>
			<li><code>data.response-text</code> - Plain text feedback provided by the user</li>
		</ul>
		<strong>Example</strong>
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
		<hr />
	</section>
	<section class="mt-12">
		<input
			type="file"
			class="file-input file-input-sm w-full shadow"
			bind:files={importedFileRefs}
		/>
		<div class="flex mt-4">
			<button class="btn btn-sm" on:click={importFeedbackFile}>Import file</button>
		</div>
	</section>
</DashLayout>
