<script lang="ts">
	import type { InputConstraints } from 'sveltekit-superforms';

	interface Option {
		label: string;
		value: string;
	}

	export let name: string;
	export let value: string;
	export let options: Option[] = [];

	export let label: string = '';
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	export let errors: any;
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	export let constraints: InputConstraints<any>[''];
	export let className: string = '';
</script>

<label class="form-control w-full">
	{#if label}
		<div class="label">
			<span class="label-text">{label}</span>
		</div>
	{/if}
	<select
		{name}
		bind:value
		class="select {errors ? 'select-error' : ''} text-sm {className} select-bordered"
		aria-invalid={errors ? 'true' : undefined}
		{...constraints}
	>
		{#each options as option}
			<option value={option.value} class="text-gray-700" selected={option.value === value}>
				{option.label}
			</option>
		{/each}
	</select>
	{#if errors}
		<div class="label">
			<span class="label-text text-xs text-red-700 dark:text-red-400">
				{errors}
			</span>
		</div>
	{/if}
</label>
