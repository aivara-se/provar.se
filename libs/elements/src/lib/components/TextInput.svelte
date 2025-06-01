<script lang="ts">
	import '../../app.css';
	import type { Snippet } from 'svelte';

	interface Props {
		name: string;
		value: string;
		label?: string;
		error?: string;
		icon?: Snippet;
		props?: object;
		className?: string;
	}

	let {
		name,
		value = $bindable(),
		label = '',
		error = '',
		icon,
		props = {},
		className = ''
	}: Props = $props();
</script>

<fieldset class="fieldset">
	{#if label}
		<legend class="fieldset-legend">{label}</legend>
	{/if}

	{#if icon}
		<label class="input {error ? 'input-error' : ''} {className}">
			{@render icon()}
			<input
				type="text"
				{name}
				bind:value
				aria-invalid={error ? 'true' : undefined}
				class="ring-0"
				{...props}
			/>
		</label>
	{:else}
		<input
			type="text"
			{name}
			bind:value
			class="input {error ? 'input-error' : ''} {className}"
			aria-invalid={error ? 'true' : undefined}
			{...props}
		/>
	{/if}

	{#if error}
		<p class="label text-error">{error}</p>
	{/if}
</fieldset>
