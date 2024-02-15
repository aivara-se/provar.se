<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import type { ActionResult } from '@sveltejs/kit';

	interface Fields {
		name: string;
	}

	let fields: Fields = {
		name: ''
	};

	function resetFields() {
		fields = {
			name: ''
		};
	}

	async function handler() {
		const data = new FormData();
		for (const [key, value] of Object.entries(fields)) {
			data.set(key, value);
		}
		const response = await fetch('', { method: 'post', body: data });
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
		resetFields();
	}
</script>

<section>
	<form class="flex flex-col gap-4" on:submit|self|stopPropagation|preventDefault={handler}>
		<label class="form-control w-full">
			<div class="label">
				<span class="label-text">Organization name</span>
			</div>
			<input type="text" id="name" class="input input-bordered w-full" bind:value={fields.name} />
		</label>
		<input type="submit" class="btn btn-md btn-block" value="Create organization" />
	</form>
</section>
