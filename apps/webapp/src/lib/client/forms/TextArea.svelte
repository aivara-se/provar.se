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

<label class="form-control">
	{#if label}
		<div class="label">
			<span class="label-text">{label}</span>
		</div>
	{/if}
	<textarea
		{name}
		bind:value
		class="textarea resize-none h-32 {errors ? 'textarea-error' : ''} {className} textarea-bordered"
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
