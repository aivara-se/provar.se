<script lang="ts">
	import '../../app.css';

	interface Option {
		label: string;
		value: string;
	}

	interface Props {
		name: string;
		value: string;
		options?: Option[];
		label?: string;
		error?: string;
		props?: object;
		className?: string;
	}

	let {
		name,
		value = $bindable(),
		options = [],
		label = '',
		error = '',
		props = {},
		className = ''
	}: Props = $props();
</script>

<fieldset class="fieldset">
	{#if label}
		<legend class="fieldset-legend">{label}</legend>
	{/if}

	<select
		{name}
		bind:value
		class="select {error ? 'select-error' : ''} {className}"
		aria-invalid={error ? 'true' : undefined}
		{...props}
	>
		{#each options as option}
			<option value={option.value} selected={option.value === value}>
				{option.label}
			</option>
		{/each}
	</select>

	{#if error}
		<p class="label text-error">{error}</p>
	{/if}
</fieldset>
