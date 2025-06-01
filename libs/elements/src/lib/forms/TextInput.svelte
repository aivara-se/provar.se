<script lang="ts">
	import type { InputConstraints } from 'sveltekit-superforms';

	export let name: string;
	export let value: string;

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
	<input
		type="text"
		{name}
		bind:value
		class="input {errors ? 'input-error' : ''} text-sm {className}"
		aria-invalid={errors ? 'true' : undefined}
		{...constraints}
	/>
	{#if errors}
		<div class="label">
			<span class="label-text text-xs text-red-400">
				{errors}
			</span>
		</div>
	{/if}
</label>
