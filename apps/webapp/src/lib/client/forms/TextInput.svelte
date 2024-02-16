<script lang="ts">
	import type { InputConstraints } from 'sveltekit-superforms';

	export let name: string;
	export let value: string;

	export let type: string = 'text';
	export let label: string = '';
	export let errors: any;
	export let constraints: InputConstraints<any>[''];
</script>

<label class="form-control w-full">
	{#if label}
		<div class="label">
			<span class="label-text">{label}</span>
		</div>
	{/if}
	<input
		type="text"
		{name}
		bind:value
		class="input {errors ? 'input-error' : ''}"
		aria-invalid={errors ? 'true' : undefined}
		{...constraints}
	/>
	{#if errors}
		<div class="label">
			<span class="label-text text-xs text-red-700 dark:text-red-400">
				{errors}
			</span>
		</div>
	{/if}
</label>
